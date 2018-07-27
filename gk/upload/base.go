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

func InitWeb() {
	DirUpload = ws.EnvParam("UploadDir")
	ImgHost = ws.EnvParam("ImgHost")
	ImgMaxWidth = ws.EnvParamInt("ImgMaxWidth", 800)
	os.MkdirAll(DirUpload, 0777)

	ws.WebGin.Static("/upload", "./upload")

	ws.WebGin.POST("/api/gk-upload/upload", func(ctx *gin.Context) {
		web := ws.Verify(ctx)
		if web.Auth {
			DoUpload(ctx, fmt.Sprint(web.SiteId))
		} else {
			ctx.Status(401)
		}
	})
	ws.WebGin.POST("/api/gk-upload/Avatar", func(ctx *gin.Context) {
		web := ws.Verify(ctx)
		if web.Auth {
			WebAvatar(ctx, fmt.Sprint(web.SiteId), web.UserId)
		} else {
			ctx.Status(401)
		}
	})

}
