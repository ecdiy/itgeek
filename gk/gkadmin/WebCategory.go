package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebTopicCategoryAdd(param *Param, res map[string]interface{}) {
	res["Id"], _ = ws.TopicCategoryDao.Add(param.SiteId, param.String("Name"), param.Int64("ParentId", 0))
	res["catList"], _ = ws.TopicCategoryDao.List(param.SiteId)
}
func WebTopicCategoryList(param *Param, res map[string]interface{}) {
	res["catList"], _ = ws.TopicCategoryDao.List(param.SiteId)
}
func WebTopicCategoryDel(param *Param, res map[string]interface{}) {
	CategoryId := param.Int64("CategoryId", 0)
	res["row"], _ = ws.TopicCategoryDao.Del(param.SiteId, CategoryId)
	ws.TopicDao.UpCategory(0, CategoryId, param.SiteId)
	res["catList"], _ = ws.TopicCategoryDao.List(param.SiteId)
}
func WebTopicCategoryModify(param *Param, res map[string]interface{}) {
	res["row"], _ = ws.TopicCategoryDao.UpName(param.String("Name"), param.Int64("Id", 0), param.SiteId)
}
