package gktopic

import (
	"github.com/gin-gonic/gin"

	"github.com/ecdiy/itgeek/gk/ws"
	"strings"
	"os"
	"io/ioutil"
)

func InitWeb(web *gin.Engine, verify func(c *gin.Context) (bool, int64)) {

	auth := func(url string, fun func(param *ws.Auth)) {
		ws.WebAuth(web, "/api/gk-topic"+url, fun, verify)
	}
	post := func(url string, fun func(param *ws.Web)) {
		ws.WebPost(web, "/api/gk-topic"+url, fun)
	}

	post("/detail", WebDetail)
	post("/topicBase", WebTopicBase)

	auth("/save", WebAdd)

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
	//--

	auth("/topicAppend", WebAppend)
	//---SEO
	web.GET("/", WebHome)
	web.GET("/p/topic/detail,:p", WebSeoDetail)
	web.GET("/sitemap.xml", WebSiteMap)

	post("/zxPage", WebZxPage)

	web.NoRoute(func(ctx *gin.Context) {
		url := ctx.Request.URL.Path
		if strings.Index(url, "/avatar/") == 0 {

			p := "./upload" + url
			_, e := os.Stat(p)
			if e == nil {
				bs, e := ioutil.ReadFile(p)
				if e == nil {
					ctx.Data(200, "image/png", bs)
					return
				}
			}
			ctx.Request.URL.Path = "/static/head-default.png"
			web.HandleContext(ctx)
			return
		}

		if strings.Index(url, "/p/") == 0 {
			ua := GetUa(ctx)
			if ua == "web" {
				ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlWeb))
			} else if ua == "h5" {
				ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlH5))
			}
			return
		}
	})
}
