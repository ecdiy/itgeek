package gknote

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebCategoryList(userId int64, param *ws.Param, res map[string]interface{}) {
	res["cat"], _ = ws.CategoryDao.List(param.SiteId, userId)
}
func WebCategoryAdd(userId int64, param *ws.Param, res map[string]interface{}) {
	res["Id"], _ = ws.CategoryDao.Add(param.SiteId, param.String("Name"), param.Int64("ParentId", 0), 0, userId)
}
func WebCategoryModify(userId int64, param *ws.Param, res map[string]interface{}) {
	res["Row"], _ = ws.CategoryDao.ModifyName(param.String("Name"), param.Int64("Id", 0), userId, param.SiteId)
}
func WebCategoryDel(userId int64, param *ws.Param, res map[string]interface{}) {

	id := param.Int64("Id", 0)
	ws.NoteDao.UpdateCat(0, id, userId, param.SiteId)
	ws.NoteDao.UpdatePCat(0, id, userId, param.SiteId)
	ws.CategoryDao.DelParentId(param.SiteId, id, userId)
	res["Row"], _ = ws.CategoryDao.Del(param.SiteId, id, userId)
}
