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
	Username := web.String("Username")
	Password := web.String("Password")
	Captcha := web.String("Captcha")
	Digits := web.String("Digits")
	if len(Username) == 0 || len(Password) == 0 {
		//web.Out["Type"] = "Password"
		web.ST(ws.StErrorParameter, "Password")
		return
	}
	v, ext, e := ws.UserDao.BaseInfoByUsername(web.SiteId, Username)
	if e != nil {
		web.ST(ws.StErrorDb)
		return
	}
	if !ext || v == nil || len(v) == 0 {
		web.ST(ws.StUsernameNotExist, "Username")
		return
	}

	errTime, _ := strconv.Atoi(v["PasswordError"])
	if errTime > 3 {
		if Captcha == "" {
			web.ST(ws.StErrorCaptcha, "Captcha", captcha.New())
			return
		} else {
			if !captcha.VerifyString(Captcha, Digits) {
				web.ST(ws.StErrorCaptcha, "Captcha", captcha.New())
				return
			}
		}
	}
	if v["Password"] == UserMd5Pass(Username, Password) {
		ws.UserDao.SetPasswordError(0, Username, web.SiteId, )
		UserInfoToRedis(web, v, )
		web.ST(ws.OK, v)
		return
	} else {
		if errTime > 3 {
			web.ST(ws.StUserPassError, "Password", captcha.New())
		} else {
			ws.UserDao.SetPasswordError(errTime+1, Username, web.SiteId, )
			web.ST(ws.StUserPassError, "Password")
		}
		return
	}
}

func UserInfoToRedis(web *ws.Web, v map[string]string) {
	delete(v, "Password")
	delete(v, "PasswordError")
	id := v["Id"]
	h := md5.New()
	h.Write([]byte(id + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))

	v["Token"] = id + "_" + tk
	id64, _ := strconv.ParseInt(id, 10, 0)
	ws.TokenDao.Del(web.SiteId, id64, web.Ua, )
	ws.TokenDao.Add(web.SiteId, id64, web.Ua, tk, )
	ws.TokenMap[web.Ua+"_"+id] = []string{tk, v["Username"]}

}

//-----

func WebUserRegister(web *ws.Web) {
	captchaId := web.String("CaptchaId")
	captchaVal := web.String("CaptchaVal")

	if !captcha.VerifyString(captchaId, captchaVal) {
		web.Result(captcha.New())
		web.ST(ws.StErrorCaptcha)
		return
	}
	Email := web.String("Email")
	Password := web.String("Password")
	Username := web.String("Username")
	Mobile := web.String("Mobile")

	if len(Username) < 5 || strings.Index(Username, "@") > 0 ||
		len(Password) < 6 || strings.Index(Email, "@") < 0 ||
		len(Mobile) < 8 || len(Mobile) > 32 {
		seelog.Warn("参数不合法.", web)
		web.Result(captcha.New())
		web.ST(ws.StErrorParameter)
		return
	}
	uCount, unb, _ := ws.UserDao.CheckByUsername(web.SiteId, Username)
	seelog.Info("~~~UserDao Register~~", Username, Mobile, Email)
	if unb && uCount > 0 {
		web.Result(captcha.New())
		web.ST(ws.StUsernameExist)
		return
	}
	eCount, eb, _ := ws.UserDao.CheckByEmail(web.SiteId, Email)
	if eb && eCount > 0 {
		web.Result(captcha.New())
		web.ST(ws.StEmailExist)
		return
	}
	Password = UserMd5Pass(Username, Password)
	id, ue := ws.UserDao.Add(web.SiteId, Username, Password, Email, Mobile)

	if ue == nil {
		v, _, _ := ws.UserDao.BaseInfo(web.SiteId, id)
		UserInfoToRedis(web, v)
		web.Out["Info"] = v
		ws.KvDao.UserCount(web.SiteId, web.SiteId)

		fee := ws.GetSoreRule(web.SiteId).Register
		if fee != 0 {
			UpCount(&ws.UpReq{UserId: id, Fee: fee, EntityId: fmt.Sprint(id), SiteId: web.SiteId,
				ScoreType: "初始资本",
				ScoreDesc: `获得初始资本 ` + fmt.Sprint(fee)})
		}

		return
	} else {
		web.Result(captcha.New())
		web.ST(ws.StErrorUnknown)
		return
	}

}
