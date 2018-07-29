package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebUserList(web *ws.Web) {
	web.Out["list"], _ = ws.UserDao.List(web.SiteId, web.Start())
	web.Out["total"], _, _ = ws.UserDao.Count(web.SiteId)
}
