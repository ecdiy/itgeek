package gknote

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebUserList(userId int64, param *ws.Param, res map[string]interface{}) {
	res["cat"], _ = ws.CategoryDao.List(param.SiteId, userId)
	res["list"], _ = ws.NoteDao.List(param.SiteId, userId)
}

func WebList(param *ws.Param, res map[string]interface{}) {
	userId := param.Int64("userId", 0)
	res["cat"], _ = ws.CategoryDao.List(param.SiteId, userId)
	res["list"], _ = ws.NoteDao.List(param.SiteId, userId)
}

func WebAdd(UserId int64, param *ws.Param, res map[string]interface{}) {
	res["Id"], _ = ws.NoteDao.Add(param.SiteId, param.String("CategoryId"), param.String("Title"), param.String("Body"), param.String("SourceType"), param.String("Source"), UserId)
}

func WebDetail(UserId int64, param *ws.Param, res map[string]interface{}) {
	res["Notes"], _, _ = ws.NoteDao.Detail(param.SiteId, param.Int64("id", 0), UserId)
}
func WebUserDetail(param *ws.Param, res map[string]interface{}) {
	res["Notes"], _, _ = ws.NoteDao.Detail(param.SiteId, param.Int64("id", 0), param.Int64("userId", 0))
}
