package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebTopicList(web *ws.Web) {

	web.Out["list"], _ = ws.TopicDao.List(web.SiteId, web.Start())
	web.Out["total"], _, _ = ws.TopicDao.Count(web.SiteId)

}
