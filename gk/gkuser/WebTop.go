package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebTopDau(web *ws.Web) {
	web.Out["list"], _ = ws.UserDao.Dau(web.SiteId)
}
