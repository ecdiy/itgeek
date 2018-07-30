package gktopic

import (
	"github.com/gin-gonic/gin"

	"github.com/ecdiy/itgeek/gk/ws"
	"strings"
	"os"
	"io/ioutil"
	"github.com/cihub/seelog"
)

func InitWeb() {

	auth := func(url string, fun func(param *ws.Web)) {

		ws.WebAuth("/api/gk-topic"+url, fun)
	}
	post := func(url string, fun func(param *ws.Web)) {
		ws.WebPost("/api/gk-topic"+url, func(web *ws.Web) {
			ws.Verify(web.Context)
			fun(web)
		})
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
	ws.WebGin.GET("/", WebHome)
	ws.WebGin.GET("/p/topic/detail,:p", WebSeoDetail)
	ws.WebGin.GET("/sitemap.xml", WebSiteMap)

	post("/zxPage", WebZxPage)

	ws.WebGin.NoRoute(func(ctx *gin.Context) {
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
			ws.WebGin.HandleContext(ctx)
			return
		}
		if strings.Index(url, "/admin/p/") == 0 {
			seelog.Error(url)
			ctx.Redirect(302, "/admin/index.html")
			return
		}
		if strings.Index(url, "/p/") == 0 {
			ua := ws.GetUa(ctx)
			if ua == "web" {
				ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlWeb))
			} else if ua == "h5" {
				ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlH5))
			}
			return
		}
	})
}
