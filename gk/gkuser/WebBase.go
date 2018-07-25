package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebMemberInfo(param *ws.Param, res map[string]interface{}) {
	username := param.String("username")
	m, bb, _ := ws.UserDao.MemberInfo(param.SiteId, username)
	if bb {
		res["member"] = m
		res["DauOrder"], _, _ = ws.UserDao.DauOrder(param.SiteId, m["Id"])
	}
}



func CountInfo(param *ws.Param, res map[string]interface{}) {
	list, _ := ws.KvDao.CountInfo(param.SiteId, )
	for _, v := range list {
		res[v[0]] = v[1]
	}
}
