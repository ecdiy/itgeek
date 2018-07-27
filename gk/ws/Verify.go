package ws

import (
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) *Web {
	auth := &Web{}
	auth.Ua, _ = c.Cookie("ua")
	if auth.Ua == "web" {
		auth.Ua = "web"
	} else {
		auth.Ua = "h5"
	}
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
				if v["Token"] == sut[idx+1:] {
					UserDao.DauAdd(auth.SiteId, uid)
					auth.UserId = uid
					auth.Username = v["Username"]
					auth.Auth = true
					return auth
				}
			} else {
				tk, b, ee := TokenDao.Find(auth.SiteId, uid, auth.Ua)
				if ee == nil && b && tk == sut[idx+1:] {
					vf, _, _ := UserDao.BaseInfo(auth.SiteId, id)
					vf["Token"] = tk
					TokenMap[auth.Ua+"_"+id] = vf
					UserDao.DauAdd(auth.SiteId, uid)
					auth.UserId = uid
					auth.Username = v["Username"]
					auth.Auth = true
					return auth
				}
			}
		}
	}
	auth.Auth = false
	return auth
}
