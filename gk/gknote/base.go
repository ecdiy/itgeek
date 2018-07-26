package gknote

import (
	"github.com/gin-gonic/gin"

	"github.com/ecdiy/itgeek/gk/ws"
)

func InitWeb(web *gin.Engine, verify func(c *gin.Context) (bool, int64)) {
	post := func(url string, fun func(param *ws.Web)) {
		ws.WebPost(web, "/api/gk-note"+url, func(param *ws.Web) {
			verify(param.Context)
			fun(param)
		})
	}
	auth := func(url string, fun func(auth *ws.Auth)) {
		ws.WebAuth(web, "/api/gk-note"+url, fun, verify)
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
