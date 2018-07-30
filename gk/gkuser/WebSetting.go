package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
	"github.com/cihub/seelog"
)

func WebSite(web *ws.Web) {
	web.Out["site"], _, _ = ws.KvDao.Get(web.SiteId, "BaseInfo")

	web.Out["SiteId"] = web.SiteId
	if web.Auth {
		mm, bx, _ := ws.UserDao.BaseInfo(web.SiteId, web.UserId)
		if bx {
			web.Out["Info"] = mm
		} else {
			seelog.Warn("not find info.", web.UserId)
		}
		web.Out["Login"] = true
	} else {
		web.Out["Login"] = false
	}
}

func WebSettingSave(auth *ws.Web) {
	auth.Out["row"], _ = ws.UserDao.Setting(auth.String("EditType"), auth.String("Info"), auth.Int64("Privacy"), auth.UserId, auth.SiteId)
}
func WebSettingGet(auth *ws.Web) {
	auth.Out["row"], _, _ = ws.UserDao.SettingGet(auth.SiteId, auth.UserId)
}

func WebSettingUpPass(auth *ws.Web) {
	p := auth.String("Password")
	if len(p) > 3 {
		p = UserMd5Pass(auth.Username, p)
		ws.UserDao.UpPassword(p, auth.UserId, auth.SiteId)
		auth.ST(ws.OK)
	} else {
		auth.ST(ws.StErrorParameter)
	}
}
