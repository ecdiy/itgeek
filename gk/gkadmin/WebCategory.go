package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebTopicCategoryAdd(web *ws.Web) {
	web.Out["Id"], _ = ws.TopicCategoryDao.Add(web.SiteId, web.String("Name"), web.Int64("ParentId" ))
	web.Out["catList"], _ = ws.TopicCategoryDao.List(web.SiteId)
}
func WebTopicCategoryList(web *ws.Web) {
	web.Out["catList"], _ = ws.TopicCategoryDao.List(web.SiteId)
}
func WebTopicCategoryDel(web *ws.Web) {
	CategoryId := web.Int64("CategoryId" )
	web.Out["row"], _ = ws.TopicCategoryDao.Del(web.SiteId, CategoryId)
	ws.TopicDao.UpCategory(0, CategoryId, web.SiteId)
	web.Out["catList"], _ = ws.TopicCategoryDao.List(web.SiteId)
}
func WebTopicCategoryModify(web *ws.Web) {
	web.Out["row"], _ = ws.TopicCategoryDao.UpName(web.String("Name"), web.Int64("Id" ), web.SiteId)
}
