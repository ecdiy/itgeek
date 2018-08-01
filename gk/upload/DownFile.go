package upload

import (
	"github.com/ecdiy/itgeek/gk/ws"
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/cihub/seelog"
	"fmt"
	"github.com/ecdiy/itgeek/gk/gkuser"
)

//Content-Length:207081
//Content-MD5:wHlX4Ow4JwWslrUuG4eA0w==
func WebResDownFile(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/octet-stream")
	auth := ws.Verify(ctx)
	if !auth.Auth {
		seelog.Error("not login,down file error.")
		return
	} else {
		idStr, _ := auth.Context.GetQuery("Id")
		id, _ := strconv.ParseInt(idStr, 10, 0)
		res, rb, _ := ws.ResourceDao.FindBySiteIdAndId(auth.SiteId, id)
		if res.UserId != auth.UserId {
			//check down log
			ext, db, _ := ws.ResourceDownDao.Exist(auth.SiteId, auth.UserId, id)
			if db && ext == 0 {
				resLink := `<go onclick="vgo('/p/down/item,` + idStr + `,` + fmt.Sprint(res.UserId) +
					`',this)">` + res.ResName + "</go>"
				duLink := `<go onclick="vgo('/p/member/` + auth.Username + `')">` + auth.Username + "</go>"
				_, _, usc, _ := gkuser.ChangeScore(auth.SiteId, "down:"+idStr, "下载资源",
					`下载了 <go onclick="vgo('/p/member/`+res.Username+`')">`+res.Username+"</go> 上传的 "+resLink,
					(-1)*res.ResScore, auth.UserId)
				if usc {
					pFee := res.Size / 1024 / 1024
					if pFee == 0 {
						pFee = 1
					}
					aFee := res.ResScore - pFee
					if aFee <= 0 {
						seelog.Error("资源积分设置错误")
						return
					}
					_, _, usc2, _ := gkuser.ChangeScore(auth.SiteId, "down:"+idStr, "下载资源", duLink+` 下载了资源 › `+resLink,
						aFee, res.UserId)
					if usc2 {
						ws.ResourceDownDao.Add(auth.SiteId, auth.UserId, id, res.ResScore, aFee, auth.Username)
					} else {
						return
					}
				}
			}
		}
		if rb {
			auth.Context.Header("Content-Disposition", "attachment;filename=\""+res.FileName+"\"")
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
