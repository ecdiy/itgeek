package gknote

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebCategoryList(auth *ws.Auth) {
	auth.Out["cat"], _ = ws.CategoryDao.List(auth.SiteId, auth.UserId)
}
func WebCategoryAdd(auth *ws.Auth) {
	auth.Out["Id"], _ = ws.CategoryDao.Add(auth.SiteId, auth.String("Name"), auth.Int64("ParentId"), 0, auth.UserId)
}
func WebCategoryModify(auth *ws.Auth) {
	auth.Out["Row"], _ = ws.CategoryDao.ModifyName(auth.String("Name"), auth.Int64("Id"), auth.UserId, auth.SiteId)
}
func WebCategoryDel(auth *ws.Auth) {
	id := auth.Int64("Id")
	ws.NoteDao.UpdateCat(0, id, auth.UserId, auth.SiteId)
	ws.NoteDao.UpdatePCat(0, id, auth.UserId, auth.SiteId)
	ws.CategoryDao.DelParentId(auth.SiteId, id, auth.UserId)
	auth.Out["Row"], _ = ws.CategoryDao.Del(auth.SiteId, id, auth.UserId)
}
