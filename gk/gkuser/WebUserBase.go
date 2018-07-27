package gkuser

import (
	"github.com/dchest/captcha"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strings"
	"github.com/cihub/seelog"
	"github.com/ecdiy/itgeek/gk/ws"
	"fmt"
)

func UserMd5Pass(Username, pass string) string {
	h := md5.New()
	h.Write([]byte( Username + "," + pass))
	return hex.EncodeToString(h.Sum(nil))
}
func WebUserLogin(web *ws.Web) {
	r := DoUserLogin(web, web.String("Username"), web.String("Password"), web.String("Captcha"), web.String("Digits"), web.Out)
	web.Out["Result"] = r.Result
	web.Out["Status"] = r.Status
}

func DoUserLogin(param *ws.Web, Username, Password, Captcha, Digits string, res map[string]interface{}) *ws.Result {
	if len(Username) == 0 || len(Password) == 0 {
		return ws.StErrorParameter.Result("Password")
	}
	v, ext, e := ws.UserDao.BaseInfoByUsername(param.SiteId, Username)
	if e != nil {
		return ws.StErrorDb.ResultNil()
	}
	if !ext {
		return ws.StUsernameNotExist.Result("Username")
	}
	if v == nil || len(v) == 0 {
		return ws.StUsernameNotExist.Result("Username")
	}
	errTime, _ := strconv.Atoi(v["PasswordError"])
	if errTime > 3 {
		if Captcha == "" {
			return ws.StErrorCaptcha.Result("Captcha").Put("Val", captcha.New())
		} else {
			if !captcha.VerifyString(Captcha, Digits) {
				return ws.StErrorCaptcha.Result("Captcha").Put("Val", captcha.New())
			}
		}
	}
	if v["Password"] == UserMd5Pass(Username, Password) {
		ws.UserDao.SetPasswordError(0, Username, param.SiteId, )
		res["Info"] = v
		return UserInfoToRedis(param.SiteId, param.Ua, v, )
	} else {
		ws.UserDao.SetPasswordError(errTime+1, Username, param.SiteId, )
		return ws.StUserPassError.Result("Password")
	}
}

func UserInfoToRedis(siteId int64, ua string, v map[string]string) *ws.Result {
	delete(v, "Password")
	delete(v, "PasswordError")

	h := md5.New()
	h.Write([]byte(v["Id"] + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))
	v["Token"] = tk

	result := ws.OK.ResultNil()
	result.Result = v["Id"] + "_" + tk
	ws.TokenDao.Del(ua, v["Id"], siteId)
	ws.TokenDao.Add(v["Id"], ua, tk, siteId)
	ws.TokenMap[ua+"_"+v["Id"]] = v
	delete(result.Param, "Token")

	return result
}

//-----

func WebUserRegister(param *ws.Web) {
	r := doUserRegister(param, param)
	param.Out["Result"] = r.Result
	param.Out["Param"] = r.Param
	param.Out["Status"] = r.Status
}

func doUserRegister(m *ws.Web, param *ws.Web) *ws.Result {
	captchaId := m.String("CaptchaId")
	captchaVal := m.String("CaptchaVal")
	if captchaId != "" && captchaVal != "" {
		if !captcha.VerifyString(captchaId, captchaVal) {
			return ws.StErrorCaptcha.Result(captcha.New())
		}
		Email := m.String("Email")
		Password := m.String("Password")
		Username := m.String("Username")
		Mobile := m.String("Mobile")

		if len(Username) < 5 || strings.Index(Username, "@") > 0 ||
			len(Password) < 6 || strings.Index(Email, "@") < 0 ||
			len(Mobile) < 8 || len(Mobile) > 32 {
			seelog.Warn("参数不合法.", m)
			return ws.StErrorParameter.Result(captcha.New())
		}
		uCount, unb, _ := ws.UserDao.CheckByUsername(m.SiteId, Username)
		seelog.Info("~~~UserDao Register~~", Username, Mobile, Email)
		if unb && uCount > 0 {
			return ws.StUsernameExist.Result(captcha.New())
		}
		eCount, eb, _ := ws.UserDao.CheckByEmail(m.SiteId, Email)
		if eb && eCount > 0 {
			return ws.StEmailExist.Result(captcha.New())
		}
		Password = UserMd5Pass(Username, Password)
		id, ue := ws.UserDao.Add(m.SiteId, Username, Password, Email, Mobile)

		if ue == nil {
			v, _, _ := ws.UserDao.BaseInfo(m.SiteId, id)
			rs := UserInfoToRedis(m.SiteId, m.Ua, v)
			param.Out["Info"] = v
			ws.KvDao.UserCount(m.SiteId, m.SiteId)

			fee := ws.GetSoreRule(m.SiteId).Register
			if fee != 0 {
				UpCount(&ws.UpReq{UserId: id, Fee: fee, EntityId: fmt.Sprint(id), SiteId: m.SiteId,
					ScoreType: "初始资本",
					ScoreDesc: `获得初始资本 ` + fmt.Sprint(fee)})
			}
			return rs
		} else {
			return ws.StErrorUnknown.Result(captcha.New())
		}
	} else {
		seelog.Warn("注册参数错误:", m)
		return ws.StErrorParameter.Result(captcha.New())
	}
}
