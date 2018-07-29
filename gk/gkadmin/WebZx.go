package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebZxSave(web *ws.Web) {
	_, ext, _ := ws.ZxDao.GetTest(web.SiteId, web.String("PageUa"), web.String("PageKey"))
	if ext {
		web.Out["row"], _ = ws.ZxDao.UpJson(web.String("Json"), web.String("PageUa"), web.String("PageKey"), web.SiteId)
	} else {
		web.Out["row"], _ = ws.ZxDao.Add(web.SiteId, web.String("PageKey"), web.String("PageUa"), web.String("Json"))
	}
}
func WebZxPub(web *ws.Web) {
	web.Out["row"], _ = ws.ZxDao.Pub(web.String("PageUa"), web.String("PageKey"), web.SiteId)
}
func WebZxGet(web *ws.Web) {
	web.Out["zxJson"], _, _ = ws.ZxDao.GetTest(web.SiteId, web.String("PageUa"), web.String("PageKey"))
}
