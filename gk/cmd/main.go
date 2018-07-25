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
)

var web = gin.New()

func ParamBase() {

	ws.ParamInit(ws.EnvDev, `
UploadDir=D:\dev\upload\
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://127.0.0.1
ImgMaxWidth=800
`)
	ws.ParamInit(ws.EnvProd, `
UploadDir=/data/upload/
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://s.ecdiy.cn
ImgMaxWidth=800
`)

}

func main() {
	ParamBase()
	defer seelog.Flush()
	ws.EnvParamSet(ws.KeyBindAddr, ":9000")

	ws.InitDao()

	gkuser.InitWeb(web)
	gknote.InitWeb(web, gkuser.Verify)
	gktopic.InitWeb(web, gkuser.Verify)
	upload.InitWeb(web, gkuser.Verify)

	gkadmin.InitWeb(web)
	web.Static("/h5dist", "./m/h5dist")
	web.Static("/dist", "./web/dist")
	web.Run(ws.EnvParam(ws.KeyBindAddr))
}
