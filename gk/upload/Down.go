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
		_, sb, _ := ws.ResourceSummaryDao.Get(auth.SiteId, auth.UserId)
		if sb {
			ws.ResourceSummaryDao.UpUploadItem(auth.UserId, auth.SiteId, auth.UserId)
		} else {
			ws.ResourceSummaryDao.Add(auth.SiteId, auth.UserId, 1, 0, 50)
		}
	}
}

func WebUploadResource(c *gin.Context) {
	web := ws.Verify(c)
	if web.Auth {
		mf, _ := c.MultipartForm()
		if len(mf.File) == 1 {
			for k, _ := range mf.File {
				up := &UpFile{}
				up.Upload(c, "", DownDir, k, fmt.Sprint(web.SiteId))
				if !up.Exist {
					id, _ := ws.ResourceDao.Add(web.SiteId, web.UserId, up.Size, up.FileName, up.Uri, web.Username)
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

type UpFile struct {
	Exist               bool
	Uri, Path, FileName string
	Size                int64
}

func (u *UpFile) Upload(c *gin.Context, extDefault, dir, key, siteId string) {
	tmp := strconv.FormatInt(time.Now().UnixNano(), 16)
	file, header, err := c.Request.FormFile(key)
	u.FileName = header.Filename
	ext := extDefault
	index := strings.Index(u.FileName, ".")
	if index > 0 {
		ext = u.FileName [ index:]
	}
	tmpFileName := DirUpload + tmp + ext
	out, err := os.Create(tmpFileName)
	if err != nil {
		seelog.Error("上传文件创建临时文件夹失败!", tmpFileName, err)
	} else {
		u.Size, _ = io.Copy(out, file)
		out.Close()
		file.Close()
		md5Name, e := Md5File(tmpFileName)
		if e == nil {
			u.Path, u.Uri = FmtImgDir(dir+siteId+"/", md5Name+ext)
			u.Uri = siteId + "/" + u.Uri
			if _, err := os.Stat(u.Path); os.IsNotExist(err) {
				u.Exist = false
				os.Rename(tmpFileName, u.Path)
			} else {
				u.Exist = true
				os.Remove(tmpFileName)
			}
		}
	}
}

func WebResourceInfo(auth *ws.Web) {
	s, sb, _ := ws.ResourceSummaryDao.Get(auth.SiteId, auth.UserId)
	if !sb {
		auth.Out["info"] = map[string]string{"UploadItem": "0", "DownItem": "0", "UpLimit": "50"}
	} else {
		auth.Out["info"] = s
		auth.Out["list"], _ = ws.ResourceDao.List(auth.SiteId, auth.UserId)
	}
}

func WebResourceList(web *ws.Web) {
	web.Out["list"], _ = ws.ResourceDao.ListTitleNew(web.SiteId, web.Start())
}

func WebResourceDetail(web *ws.Web) {
	web.Out["detail"], _, _ = ws.ResourceDao.GetBySiteIdAndId(web.SiteId, web.Int64("UserId"), web.Int64("Id"))
}
