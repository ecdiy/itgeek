package main

import (
	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
	"github.com/ecdiy/itgeek/gk/gkuser"

	"github.com/ecdiy/itgeek/gk/gktopic"
	"github.com/ecdiy/itgeek/gk/gknote"
	"github.com/ecdiy/itgeek/gk/upload"
	"github.com/ecdiy/itgeek/gk/ws"
	"github.com/ecdiy/itgeek/gk/gkadmin"
	"os"
	"strings"
)

var web = gin.New()

func ParamBase() {

	ws.ParamInit(ws.EnvDev, `
UploadDir=./upload/
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://127.0.0.1:9000
ImgMaxWidth=800
MultiSite=0
`)
	ws.ParamInit(ws.EnvProd, `
UploadDir=./upload/
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://s.ecdiy.cn
ImgMaxWidth=800
MultiSite=0
`)

}

var BaseDir string

func init() {
	bin := os.Args[0]
	idx := strings.LastIndex(bin, string(os.PathSeparator))
	if idx > 0 {
		BaseDir = bin[0 : idx+1]
	} else {
		BaseDir, _ = os.Getwd()
	}
}

func main() {

	ParamBase()
	defer seelog.Flush()

	ws.EnvParamSet(ws.KeyBindAddr, ":9000")
	ws.MultiSite = ws.EnvParamInt("MultiSite", 0)

	ws.InitDao()

	gkuser.InitWeb(web)

	gknote.InitWeb(web, gkuser.Verify)
	gktopic.InitWeb(web, gkuser.Verify)
	upload.InitWeb(web, gkuser.Verify)

	gkadmin.InitWeb(web)
	web.Static("/h5dist", BaseDir+"/m/h5dist")
	web.Static("/dist", BaseDir+"/web/dist")
	web.Static("/static", BaseDir+"/static")

	seelog.Info("version:0.0.1 (itgeek.top) BaseDir=", BaseDir)
	web.Run(ws.EnvParam(ws.KeyBindAddr))
}
