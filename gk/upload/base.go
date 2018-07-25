package upload

import (
	"github.com/gin-gonic/gin"

	"os"
	"github.com/ecdiy/itgeek/gk/ws"
	"fmt"
)

const (
	UriUpload = "/upload/"
)

var (
	DirUpload   string
	ImgHost     string
	ImgMaxWidth int
)

func InitWeb(web *gin.Engine, verify func(c *gin.Context) (bool, int64)) {
	DirUpload = ws.EnvParam("UploadDir")
	ImgHost =  ws.EnvParam("ImgHost")
	ImgMaxWidth = ws.EnvParamInt("ImgMaxWidth", 800)
	os.MkdirAll(DirUpload, 0777)
	web.POST("/api/gk-upload/upload", func(ctx *gin.Context) {
		b, _ := verify(ctx)
		if b {
			DoUpload(ctx, fmt.Sprint(ws.SiteId(ctx)))
		} else {
			ctx.Status(401)
		}
	})
	web.POST("/api/gk-upload/Avatar", func(ctx *gin.Context) {
		b, userId := verify(ctx)
		if b {
			WebAvatar(ctx, fmt.Sprint(ws.SiteId(ctx)), userId)
		} else {
			ctx.Status(401)
		}
	})

}
