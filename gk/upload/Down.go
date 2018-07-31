package upload

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"github.com/ecdiy/itgeek/gk/ws"
	"strconv"
	"time"
	"strings"
	"os"
	"io"
	"github.com/cihub/seelog"
)

func WebPublicResource(auth *ws.Web) {
	rn := auth.String("ResName")
	rd := auth.String("ResDesc")
	if len(rn) > 3 && len(rd) > 10 {
		sc := auth.Int64("ResScore")
		auth.Out["row"], _ = ws.ResourceDao.Up(auth.String("ResName"), auth.String("ResDesc"), sc, auth.Int64("Id"), auth.UserId)
	}
}

func WebUploadResource(c *gin.Context) {
	web := ws.Verify(c)
	if web.Auth {
		mf, _ := c.MultipartForm()
		if len(mf.File) == 1 {
			for k, _ := range mf.File {
				ex, path, _, size := doUploadFileMd5Res(c, "", DownDir, k, fmt.Sprint(web.SiteId))
				if ex {
					id, _ := ws.ResourceDao.Add(web.SiteId, web.UserId, size, path)
					web.Out["Id"] = id
				} else {
					web.Out["Id"] = 0
				}
			}
		}
		c.JSON(200, web.Out)
	} else {
		c.Status(401)
	}
}

func doUploadFileMd5Res(c *gin.Context, extDefault, dir, key, siteId string) (bool, string, string, int64) {
	tmp := strconv.FormatInt(time.Now().UnixNano(), 16)
	file, header, err := c.Request.FormFile(key)

	filename := header.Filename
	ext := extDefault
	index := strings.Index(filename, ".")
	if index > 0 {
		ext = filename[ index:]
	}
	tmpFileName := DirUpload + tmp + ext
	out, err := os.Create(tmpFileName)

	if err != nil {
		seelog.Error("上传文件创建临时文件夹失败!", tmpFileName, err)
	} else {
		size, err2 := io.Copy(out, file)
		out.Close()
		file.Close()
		if err2 != nil {
			seelog.Error("写入上传文件流时失败!", tmpFileName, err)
			return false, "", "", size
		}
		md5Name, e := Md5File(tmpFileName)
		if e == nil {
			pre, uri := FmtImgDir(dir+siteId+"/", md5Name)
			path := pre + ext
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Rename(tmpFileName, path)
				return true, path, uri, size
			} else {
				os.Remove(tmpFileName)
				return false, path, uri, size
			}
		}
	}
	return false, "", "", 0
}

func WebResourceInfo(auth *ws.Web) {
	s, sb, _ := ws.ResourceSummaryDao.Get(auth.SiteId, auth.UserId)
	if !sb {
		auth.Out["info"] = map[string]string{"UploadItem": "0", "DownItem": "0", "UpLimit": "50"}
	} else {
		auth.Out["info"] = s
	}
}
