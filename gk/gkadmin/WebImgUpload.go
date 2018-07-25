package gkadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/ecdiy/itgeek/gk/upload"
)

func WebUpload(context *gin.Context) {

	upload.DoUpload(context, context.PostForm("SiteId"))

}
