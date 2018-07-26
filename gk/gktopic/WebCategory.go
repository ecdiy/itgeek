package gktopic

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebCategoryList(web *ws.Web) {
	web.Out["catList"], _ = ws.TopicCategoryDao.List(web.SiteId)
}
