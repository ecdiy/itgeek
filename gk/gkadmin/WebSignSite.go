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
	c.JSON(200, auth.Out)
}
func WebNewAdminUser(auth *ws.Web) {
	us := auth.String("SiteAdminUser")
	ps := auth.String("SiteAdminPass")
	if len(us) > 3 && len(ps) > 3 {
		h := md5.New()
		h.Write([]byte( us + "," + ps))
		ws.KvDao.Update(us, "SiteAdminUser", auth.SiteId)
		ws.KvDao.Update(hex.EncodeToString(h.Sum(nil)), "SiteAdminPass", auth.SiteId)
	}
	auth.ST(ws.OK)
}
func WebAdminUserInit(c *gin.Context) {
	auth := ws.VerifyAdmin(c)
	_, u, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminUser")
	_, p, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminPass")
	if !u && !p {
		us := auth.String("SiteAdminUser")
		ps := auth.String("SiteAdminPass")
		if len(us) > 3 && len(ps) > 3 {
			h := md5.New()
			h.Write([]byte( us + "," + ps))
			ws.KvDao.Add(0, "SiteAdminUser", us)
			ws.KvDao.Add(0, "SiteAdminPass", hex.EncodeToString(h.Sum(nil)))
			userLogin(auth)
		}
	} else {
		auth.ST(ws.StExist)
		c.JSON(200, auth.Out)
	}
}

func userLogin(auth *ws.Web) {
	us := auth.String("SiteAdminUser")
	_, u, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminUser")
	md5pass, p, _ := ws.KvDao.Get(auth.SiteId, "SiteAdminPass")
	if u && p {
		h := md5.New()
		h.Write([]byte( us + "," + auth.String("SiteAdminPass")))
		ep := hex.EncodeToString(h.Sum(nil))
		if ep == md5pass {
			tk := Token(us)
			ua := ws.GetUa(auth.Context)
			ws.KeySave(auth.SiteId, ua+ws.KvGeekAdmin, tk)
			auth.ST(ws.OK, tk)
			auth.Context.JSON(200, auth.Out)
			return
		}
	}
	auth.ST(ws.StUserPassError)
	auth.Context.JSON(200, auth.Out)
}

func WebAdminUserLogin(c *gin.Context) {
	auth := ws.VerifyAdmin(c)
	userLogin(auth)
}

func Token(id string) string {
	h := md5.New()
	h.Write([]byte(id + ";" + strconv.FormatInt(time.Now().UnixNano(), 16)))
	tk := hex.EncodeToString(h.Sum(nil))
	return tk
}
