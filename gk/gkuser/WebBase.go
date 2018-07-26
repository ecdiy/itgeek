package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebMemberInfo(web *ws.Web) {
	username := web.String("username")
	m, bb, _ := ws.UserDao.MemberInfo(web.SiteId, username)
	if bb {
		web.Out["member"] = m
		web.Out["DauOrder"], _, _ = ws.UserDao.DauOrder(web.SiteId, m["Id"])
	}
}

func CountInfo(web *ws.Web) {
	list, _ := ws.KvDao.CountInfo(web.SiteId, )
	for _, v := range list {
		web.Out[v[0]] = v[1]
	}
}
