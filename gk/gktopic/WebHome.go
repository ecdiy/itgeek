package gktopic

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/ecdiy/itgeek/gk/ws"
)

var (
	HtmlWeb = ""
	HtmlH5  = ""
)

func init() {
	w, e := ioutil.ReadFile("web/index.html")
	if e == nil {
		HtmlWeb = string(w)
	}
	w2, e2 := ioutil.ReadFile("m/index.html")
	if e2 == nil {
		HtmlH5 = string(w2)
	}
}
func WebHome(ctx *gin.Context) {
	ua := ws.GetUa(ctx)
	if ua == "web" {
		ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlWeb))
	} else if ua == "h5" {
		ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlH5))
	} else {
		WebSeoHome(ctx)
	}
}
