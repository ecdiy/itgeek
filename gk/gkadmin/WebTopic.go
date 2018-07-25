package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebTopicList(param *Param, res map[string]interface{}) {

	res["list"], _ = ws.TopicDao.List(param.SiteId, param.Start(20))
	res["total"], _, _ = ws.TopicDao.Count(param.SiteId)

}
