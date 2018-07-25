package upload

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
	"strconv"
	"strings"
	log "github.com/cihub/seelog"
	"encoding/hex"
	"io/ioutil"
	"crypto/md5"
)

func FmtImgDir(dir, uri string) (string, string) {
	ln := len(uri)
	if ln > 6 {
		x := uri[0:3] + "/" + uri[3:6]
		os.MkdirAll(dir+x, 0644)
		urx := x + "/" + uri[6:]
		path := dir + urx
		return path, urx
	}
	if ln > 3 {
		x := uri[0:3]
		os.MkdirAll(dir+x, 0644)
		urx := x + "/" + uri[3:]
		path := dir + urx
		return path, urx
	}
	return dir + uri, uri
}
func Md5Byte(bs []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(bs)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func Md5File(fileName string) (string, error) {
	bs, e := ioutil.ReadFile(fileName)
	if e == nil {
		return Md5Byte(bs), nil
	} else {
		return "", e
	}
}

func WebAvatar(ctx *gin.Context, siteId string, userId int64) {
	mf, _ := ctx.MultipartForm()
	if len(mf.File) == 1 {
		for k, _ := range mf.File {
			dir := DirUpload + "avatar/" + siteId + "/" + strconv.FormatInt(userId, 10)
			doUploadFile(ctx, dir, k)
			to := dir + "/48.png"
			ImgResize(dir+"/init", to, 48)
		}
	}
}

//---------
func DoUpload(c *gin.Context, siteId string) string {
	mf, _ := c.MultipartForm()
	res := ""
	if len(mf.File) == 1 {
		for k, _ := range mf.File {
			res += doUploadFileMd5(c, k, siteId)
		}
	}
	return res
}

func doUploadFileMd5(c *gin.Context, key, siteId string) string {
	tmp := strconv.FormatInt(time.Now().UnixNano(), 16)
	file, header, err := c.Request.FormFile(key)

	filename := header.Filename
	ext := ".png"
	index := strings.Index(filename, ".")
	if index > 0 {
		ext = filename[ index:]
	}
	tmpFileName := DirUpload + tmp + ext
	out, err := os.Create(tmpFileName)

	if err != nil {
		log.Error("上传文件创建临时文件夹失败!", tmpFileName, err)
	} else {
		_, err2 := io.Copy(out, file)
		out.Close()
		file.Close()
		if err2 != nil {
			log.Error("写入上传文件流时失败!", tmpFileName, err)
			return ""
		}
		md5Name, e := Md5File(tmpFileName)
		if e == nil {
			pre, uri := FmtImgDir(DirUpload+siteId+"/", md5Name)
			path := pre + ext
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Rename(tmpFileName, path)
			} else {
				os.Remove(tmpFileName)
			}
			if strings.ToLower(ext) != ".gif" {
				ext8 := "_" + strconv.Itoa(ImgMaxWidth) + ext
				urx := UriUpload + siteId + "/" + uri + ext8
				if _, err := os.Stat(pre + ext8); os.IsNotExist(err) {
					dx, _ := ImgResize(path, pre+ext8, ImgMaxWidth)
					if dx > ImgMaxWidth {
						c.JSON(200, map[string]interface{}{"location": ImgHost + urx})
						return path
					}
				} else {
					c.JSON(200, map[string]interface{}{"location": ImgHost + urx})
					return path
				}
			}
			urx := UriUpload + siteId + "/" + uri + ext
			c.JSON(200, map[string]interface{}{"location": ImgHost + urx})
			return path
		}
	}
	return ""
}

func doUploadFile(c *gin.Context, dir, key string) {
	os.MkdirAll(dir, 0777)
	os.Remove(dir + "/init")
	file, _, err := c.Request.FormFile(key)
	out, err := os.Create(dir + "/init")
	_, err2 := io.Copy(out, file)
	out.Close()
	file.Close()
	if err2 != nil {
		log.Error("写入上传文件流时失败!", err)
	}
	c.JSON(200, map[string]interface{}{"location": ""})
}
