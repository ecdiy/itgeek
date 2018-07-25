package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
	"github.com/cihub/seelog"
)

func WebSite(param *ws.Param, res map[string]interface{}) {
	res["site"], _, _ = ws.KvDao.Get(param.SiteId, "BaseInfo")
	isLogin, uId := Verify(param.Context)
	res["SiteId"] = param.SiteId
	if isLogin {
		mm, bx, _ := ws.UserDao.BaseInfo(param.SiteId, uId)
		if bx {
			res["Info"] = mm
		} else {
			seelog.Warn("not find info.", uId)
		}
		res["Login"] = true
	} else {
		res["Login"] = false
	}
}

func WebSettingSave(userId int64, param *ws.Param, res map[string]interface{}) {
	res["row"], _ = ws.UserDao.Setting(param.String("EditType"), param.String("Info"), param.Int64("Privacy", 0), userId, param.SiteId)
}
func WebSettingGet(userId int64, param *ws.Param, res map[string]interface{}) {
	res["row"], _, _ = ws.UserDao.SettingGet(param.SiteId, userId)
}
