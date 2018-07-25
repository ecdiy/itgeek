package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebTopDau(param *ws.Param, res map[string]interface{}) {
	res["list"], _ = ws.UserDao.Dau(param.SiteId)
}
