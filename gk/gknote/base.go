package gknote

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func InitWeb() {
	post := func(url string, fun func(param *ws.Web)) {
		ws.WebPost("/api/gk-note"+url, func(param *ws.Web) {
			ws.Verify(param.Context)
			fun(param)
		})
	}
	auth := func(url string, fun func(auth *ws.Web)) {
		ws.WebAuth("/api/gk-note"+url, fun)
	}
	post("/user/list", WebList)
	auth("/list", WebUserList)

	auth("/add", WebAdd)

	post("/user/detail", WebUserDetail)
	auth("/detail", WebDetail)

	auth("/categoryList", WebCategoryList)
	auth("/categoryAdd", WebCategoryAdd)
	auth("/categoryModify", WebCategoryModify)
	auth("/categoryDel", WebCategoryDel)

}
