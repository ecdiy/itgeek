package gktopic

import (
	"github.com/gin-gonic/gin"

	"github.com/ecdiy/itgeek/gk/ws"
	"strings"
	"os"
	"io/ioutil"
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
