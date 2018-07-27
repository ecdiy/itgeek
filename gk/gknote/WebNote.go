package gknote

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebUserList(auth *ws.Web) {
	auth.Out["cat"], _ = ws.CategoryDao.List(auth.SiteId, auth.UserId)
	auth.Out["list"], _ = ws.NoteDao.List(auth.SiteId, auth.UserId)
}

func WebList(web *ws.Web) {
	userId := web.Int64("userId")
	web.Out["cat"], _ = ws.CategoryDao.List(web.SiteId, userId)
	web.Out["list"], _ = ws.NoteDao.List(web.SiteId, userId)
}

func WebAdd(auth *ws.Web) {
	auth.Out["Id"], _ = ws.NoteDao.Add(auth.SiteId, auth.String("CategoryId"), auth.String("Title"), auth.String("Body"),
		auth.String("SourceType"), auth.String("Source"), auth.UserId)
}

func WebDetail(auth *ws.Web) {
	auth.Out["Notes"], _, _ = ws.NoteDao.Detail(auth.SiteId, auth.Int64("id"), auth.UserId)
}
func WebUserDetail(web *ws.Web) {
	web.Out["Notes"], _, _ = ws.NoteDao.Detail(web.SiteId, web.Int64("id"), web.Int64("userId"))
}
