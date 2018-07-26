package gkuser

import (
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"strconv"
	"time"
	"encoding/hex"
	"github.com/ecdiy/itgeek/gk/ws"
)

func Verify(c *gin.Context) (bool, int64) {
	sut, e := c.Cookie("token")
	ua, ep := c.Cookie("ua")
	if ep == nil && e == nil && len(sut) > 1 {
		ub, uId, un := doVerify(sut, ua, ws.SiteId(c))
		if ub {
			c.Set("UserId", uId)
			c.Set("Username", un)
		}
		return ub, uId
	}
	return false, 0
}

func Token(id int64) string {
	h := md5.New()
	ids := strconv.FormatInt(id, 10)
	h.Write([]byte(ids + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))
	return ids + "_" + tk
}

func InitWeb(web *gin.Engine) {

	post := func(url string, fun func(web *ws.Web)) {
		ws.WebPost(web, "/api/gk-user"+url, func(web *ws.Web) {
			Verify(web.Context)
			fun(web)
		})
	}
	auth := func(url string, fun func(auth *ws.Auth)) {
		ws.WebAuth(web, "/api/gk-user"+url, fun, Verify)
	}
	//xgin.SpAjax(base.IsDevEnv(), "/sp", db, web, UserFilter, "Sp")

	post("/site", WebSite)

	web.GET("/api/gk-user/Captcha", Captcha)
	web.GET("/api/gk-user/CaptchaNew", CaptchaNew)
	web.POST("/api/gk-user/CaptchaNew", CaptchaNew)

	post("/Register", WebUserRegister)
	post("/Login", WebUserLogin)

	post("/CountInfo", CountInfo)

	post("/memberInfo", WebMemberInfo)
	auth("/setting/save", WebSettingSave)
	auth("/setting/get", WebSettingGet)

	auth("/LoginAwardStatus", WebScoreLoginAward)
	auth("/LoginAwardDo", WebScoreLoginAwardDo)
	auth("/scoreLogList", WebScoreLogList)

	auth("/msgList", WebMsgList)
	auth("/msgDel", WebMsgDel)
	auth("/msgRead", WebMsgRead)

	post("/topDau", WebTopDau)
}
