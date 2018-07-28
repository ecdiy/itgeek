package gkuser

import (
	"crypto/md5"
	"strconv"
	"time"
	"encoding/hex"
	"github.com/ecdiy/itgeek/gk/ws"
)

func Token(id int64) string {
	h := md5.New()
	ids := strconv.FormatInt(id, 10)
	h.Write([]byte(ids + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))
	return ids + "_" + tk
}

func InitWeb() {

	post := func(url string, fun func(web *ws.Web)) {
		ws.WebPost("/api/gk-user"+url, func(web *ws.Web) {
			ws.Verify(web.Context)
			fun(web)
		})
	}
	auth := func(url string, fun func(auth *ws.Web)) {
		ws.WebAuth("/api/gk-user"+url, fun)
	}
	//xgin.SpAjax(base.IsDevEnv(), "/sp", db, web, UserFilter, "Sp")

	post("/site", WebSite)

	ws.WebGin.GET("/api/gk-user/Captcha", Captcha)

	post("/CaptchaNew", CaptchaNew)

	post("/Register", WebUserRegister)
	post("/Login", WebUserLogin)

	post("/CountInfo", CountInfo)

	post("/memberInfo", WebMemberInfo)
	auth("/setting/save", WebSettingSave)
	auth("/setting/get", WebSettingGet)

	auth("/LoginAwardStatus", WebScoreLoginAwardStatus)
	auth("/LoginAwardDo", WebScoreLoginAwardDo)
	auth("/scoreLogList", WebScoreLogList)

	auth("/msgList", WebMsgList)
	auth("/msgDel", WebMsgDel)
	auth("/msgRead", WebMsgRead)

	post("/topDau", WebTopDau)
}
