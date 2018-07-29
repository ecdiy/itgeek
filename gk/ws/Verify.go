package ws

import (
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
)

var (
	ecTokenMap = make(map[string]string)
	//authMap = make(map[string]func(userId int64, param *Param, res map[string]interface{}))
)

func Verify(c *gin.Context) *Web {
	auth := &Web{}
	auth.Ua = GetUa(c)
	auth.Context = c
	auth.Out = make(map[string]interface{})
	auth.SiteId = SiteId(c)

	sut, e := c.Cookie(auth.Ua + "Token")
	if e == nil && len(sut) > 1 {
		idx := strings.Index(sut, "_")
		if idx > 0 {
			id := sut[0:idx]
			v, b := TokenMap [auth.Ua+"_"+id]
			uid, _ := strconv.ParseInt(id, 10, 0)
			if b {
				if v[0] == sut[idx+1:] {
					UserDao.DauAdd(auth.SiteId, uid)
					auth.UserId = uid
					auth.Username = v[1]
					auth.Auth = true
					return auth
				}
			} else {
				tk, b, ee := TokenDao.Find(auth.SiteId, uid, auth.Ua)
				if ee == nil && b && tk == sut[idx+1:] {
					vf, _, _ := UserDao.BaseInfo(auth.SiteId, id)
					auth.Username = vf["Username"]
					TokenMap[auth.Ua+"_"+id] = []string{tk, auth.Username}
					UserDao.DauAdd(auth.SiteId, uid)
					auth.UserId = uid
					auth.Auth = true
					return auth
				}
			}
		}
	}
	auth.Auth = false
	return auth
}

func VerifyAdmin(c *gin.Context) (*Web) {
	auth := &Web{}
	auth.Context = c
	auth.Out = make(map[string]interface{})
	ua := GetUa(c)
	tKey := ua + KvGeekAdmin
	ect, _ := c.Cookie(tKey)
	auth.SiteId = 0
	auth.initParam()
	kv, kvb, _ := KvDao.Get(auth.SiteId, tKey)
	if kvb && kv == ect {
		auth.Auth = true
		return auth
	}
	auth.Auth = false
	return auth
}

func VerifyMultiSiteAdmin(c *gin.Context) (*Web) {
	auth := &Web{}
	auth.Context = c
	auth.Out = make(map[string]interface{})
	ua := GetUa(c)
	ect, e := c.Cookie("ecToken")

	if e == nil && len(ect) > 2 {
		idx := strings.Index(ect, "_")
		if idx > 0 {
			id := ect[0:idx]
			etk := ua + ect[0:idx]
			v, vb := ecTokenMap[etk]
			if !vb || v != ect[idx+1:] {
				qTk, qb, _ := Gpa.QueryString("select Token from site.EcToken where Ua=? and UserId=?", ua, id)
				if !qb {
					auth.Auth = false
					return auth
				} else {
					if qTk == ect[idx+1:] {
						ecTokenMap[etk] = qTk
					} else {
						auth.Auth = false
						return auth
					}
				}
			}
			auth.initParam()
			siteId := auth.Int64("siteId")
			if siteId > 0 {
				ext, _, _ := Gpa.QueryInt("select count(*) from site.Site where UserId=? and Id=?", id, siteId)
				if ext == 1 {
					auth.SiteId = siteId
					auth.Auth = true
					return auth
				}
			}
		}
	}

	auth.Auth = false
	return auth
}
