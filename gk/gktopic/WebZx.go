package gktopic

import "github.com/ecdiy/itgeek/gk/ws"

func WebZxPage(web *ws.Web) {
	web.Out["result"], _, _ = ws.ZxDao.Get(web.SiteId, web.Ua, web.String("PageKey"))
}
