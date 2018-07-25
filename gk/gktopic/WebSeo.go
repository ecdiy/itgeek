package gktopic

import (
	"github.com/gin-gonic/gin"
	"text/template"
	"github.com/cihub/seelog"
	"strings"
	"github.com/ecdiy/itgeek/gk/ws"
)

var (
	tplHome, tplDetail, tplSiteMap template.Template
)

func tmpHome() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>ITGeek 首页</title>
	<meta name="keywords" content="Go,Golang,Go语言">
	<meta name="description" content="Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
  </head>
<ul>
{{range .list}}<li><a href="/p/topic/detail,{{.Id}},{{.UserId}}">{{.Title}}</a></li>{{end}}
</ul>
共有{{.total}}条记录</html>
`
}

func tmpDetail() string {
	return `<!DOCTYPE html><meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>{{.detail.Title}}</title>
<body>
<div>{{.detail.Body}}</div>
<ul>
{{range .list}}{{.Reply}}{{end}}
</ul>
<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?efe65b9fdb88a6e2a8d8c94d15bdd2af";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>

</body></html>
`
}

func tmpSiteMap() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<urlset> 
{{range .list}}
<url>
<loc>http://itgeek.top/p/topic/detail,{{.Id}},{{.UserId}}</loc>
<priority>1.0</priority>
<lastmod>{{.CreateAt}}</lastmod>
<changefreq>Always</changefreq>
</url>
{{end}}
</urlset> 
`
}

func init() {
	tplHome = template.Template{}
	tplDetail = template.Template{}
	tplSiteMap = template.Template{}
	_, e := tplHome.Parse(tmpHome())
	if e != nil {
		seelog.Error("template home fail:", e)
	}
	_, e2 := tplDetail.Parse(tmpDetail())
	if e2 != nil {
		seelog.Error("template detail fail:", e)
	}
	_, e3 := tplSiteMap.Parse(tmpSiteMap())
	if e3 != nil {
		seelog.Error("template detail fail:", e)
	}
}

func WebSeoHome(ctx *gin.Context) {
	data := make(map[string]interface{})
	SiteId := xgin.SiteId(ctx)
	data["list"], _ = ws.TopicDao.Top1000(SiteId, )
	data["total"], _, _ = ws.TopicDao.Count(SiteId, )
	e := tplHome.Execute(ctx.Writer, data)
	if e != nil {
		seelog.Error(e)
	}
}

func WebSeoDetail(ctx *gin.Context) {
	SiteId := xgin.SiteId(ctx)
	p := ctx.Param("p")
	pp := strings.Split(p, ",")
	if len(pp) == 2 {
		d, b, _ := ws.TopicDao.FindById(SiteId, pp[0], pp[1])
		if b {
			data := make(map[string]interface{})
			data["detail"] = d
			data["list"], _ = ws.ReplyDao.List(SiteId, pp[0])

			e := tplDetail.Execute(ctx.Writer, data)
			if e != nil {
				seelog.Error(e)
			}
		} else {
			seelog.Error("TopicDao not exist,", pp)
		}
	} else {
		seelog.Error("param error.", p)
	}
}

func WebSiteMap(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["list"], _ = ws.TopicDao.Top1000(xgin.SiteId(ctx))
	e := tplSiteMap.Execute(ctx.Writer, data)
	if e != nil {
		seelog.Error(e)
	}
}
