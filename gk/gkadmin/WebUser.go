package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebUserList(param *Param, res map[string]interface{}) {
	res["list"], _ = ws.UserDao.List(param.SiteId, param.Start(20))
	res["total"], _, _ = ws.UserDao.Count(param.SiteId)
}
