package gktopic

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebCategoryList(param *ws.Param, res map[string]interface{}) {
	res["catList"], _ = ws.TopicCategoryDao.List(param.SiteId)
}
