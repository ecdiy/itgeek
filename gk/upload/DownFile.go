package upload

import (
	"github.com/ecdiy/itgeek/gk/ws"
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/cihub/seelog"
)

//Content-Length:207081
//Content-MD5:wHlX4Ow4JwWslrUuG4eA0w==
func WebResDownFile(ctx *gin.Context) {
	auth := ws.Verify(ctx)
	if !auth.Auth {
		seelog.Error("not login,down file error.")
		return
	} else {
		idStr, _ := auth.Context.GetQuery("Id")
		id, _ := strconv.ParseInt(idStr, 10, 0)

		res, rb, _ := ws.ResourceDao.FindBySiteIdAndId(auth.SiteId, id)
		if rb {
			auth.Context.Header("Content-Disposition", "attachment;filename=\""+res.FileName+"\"")
			auth.Context.Header("Content-Type", "application/octet-stream")
			file, err := os.Open(DownDir + res.ResPath)
			defer file.Close()
			if err == nil {
				io.Copy(auth.Context.Writer, file)
			}
		} else {
			seelog.Error("resource not exit.")
		}
	}
}
