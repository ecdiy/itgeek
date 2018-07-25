package gktopic

import (
	"github.com/gin-gonic/gin"

	"github.com/ecdiy/itgeek/gk/ws"
)

func InitWeb(web *gin.Engine, verify func(c *gin.Context) (bool, int64)) {

	post := func(url string, fun func(param *ws.Param, res map[string]interface{})) {
		ws.Post(web, "/api/gk-topic"+url, func(param *ws.Param, res map[string]interface{}) {
			verify(param.Context)
			fun(param, res)
		})
	}
	auth := func(url string, fun func(UserId int64, param *ws.Param, res map[string]interface{})) {
		ws.Auth(web, "/api/gk-topic"+url, fun, verify)
	}

	auth("/save", WebAdd)
	post("/detail", WebDetail)

	post("/topic/hot", WebTopicHot)

	post("/list", WebList)
	post("/listByUserId", WebListByUserId)
	post("/category/list", WebCategoryList)
	auth("/fav", WebFav)
	auth("/favList", WebFavList)
	auth("/favStatus", WebFavStatus)

	auth("/reply", WebReply)
	auth("/topicReplyThank", WebTopicReplyThank)

	auth("/follow", WebFollow)
	auth("/followList", WebFollowList)
	auth("/followStatus", WebFollowStatus)

	//---SEO
	web.GET("/", WebHome)
	web.GET("/p/topic/detail,:p", WebSeoDetail)
	web.GET("/sitemap.xml", WebSiteMap)

	post("/zxPage", WebZxPage)
}
