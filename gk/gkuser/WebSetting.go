package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
	"github.com/cihub/seelog"
)

func WebSite(web *ws.Web) {
	web.Out["site"], _, _ = ws.KvDao.Get(web.SiteId, "BaseInfo")
	isLogin, uId := Verify(web.Context)
	web.Out["SiteId"] = web.SiteId
	if isLogin {
		mm, bx, _ := ws.UserDao.BaseInfo(web.SiteId, uId)
		if bx {
			web.Out["Info"] = mm
		} else {
			seelog.Warn("not find info.", uId)
		}
		web.Out["Login"] = true
	} else {
		web.Out["Login"] = false
	}
}

func WebSettingSave(auth *ws.Auth) {
	auth.Out["row"], _ = ws.UserDao.Setting(auth.String("EditType"), auth.String("Info"), auth.Int64("Privacy"), auth.UserId, auth.SiteId)
}
func WebSettingGet(auth *ws.Auth) {
	auth.Out["row"], _, _ = ws.UserDao.SettingGet(auth.SiteId, auth.UserId)
}
