package gkadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/ecdiy/itgeek/gk/ws"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func WebUserStatus(c *gin.Context) {
	auth := ws.VerifyAdmin(c)
	if auth.Auth {
		auth.Out["login"] = true
	} else {
		auth.Out["login"] = false
		_, auth.Out["user"], _ = ws.KvDao.Get(auth.SiteId, "SiteAdminUser")
		_, auth.Out["pass"], _ = ws.KvDao.Get(auth.SiteId, "SiteAdminPass")
	}
}

func WebAdminUserInit(c *gin.Context) {
	auth := ws.VerifyAdmin(c)
	_, u, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminUser")
	_, p, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminPass")
	if !u && !p {
		us := auth.String("SiteAdminUser")
		ws.KvDao.Add(0, "SiteAdminUser", us)
		h := md5.New()
		h.Write([]byte( us + "," + auth.String("SiteAdminPass")))
		ws.KvDao.Add(0, "SiteAdminPass", hex.EncodeToString(h.Sum(nil)))
		auth.Result(ws.OK)
	} else {
		auth.Result(ws.StExist)
	}
}

func WebAdminUserLogin(c *gin.Context) {
	auth := ws.VerifyAdmin(c)
	us := auth.String("SiteAdminUser")
	_, u, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminUser")
	md5pass, p, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminPass")
	if u && p {
		h := md5.New()
		h.Write([]byte( us + "," + auth.String("SiteAdminPass")))
		ep := hex.EncodeToString(h.Sum(nil))
		if ep == md5pass {
			tk := Token(us)
			ua := ws.GetUa(c)
			ws.KeySave(auth.SiteId, ua+ws.KV_GeekAdmin, tk)
			auth.Result(tk)
			return
		}
	}
	auth.ST(ws.StUserPassError)
}

func Token(id string) string {
	h := md5.New()
	h.Write([]byte(id + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))
	return tk
}
