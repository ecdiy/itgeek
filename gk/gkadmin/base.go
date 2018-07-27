package gkadmin

import (
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/ecdiy/itgeek/gk/ws"
)

var ecTokenMap = make(map[string]string)

func InitWeb() {

	ws.WebGin.POST("/api/gk-admin/siteUpload", WebUpload)

	auth := func(url string, fun func(param *Param, res map[string]interface{})) {
		ws.WebGin.POST("/api/gk-admin/"+url, func(c *gin.Context) {
			auth, param := VerifyAdmin(c)
			if auth {
				res := make(map[string]interface{})
				fun(param, res)
				c.JSON(200, res)
			} else {
				c.AbortWithStatus(401)
			}
		})
	}

	auth("baseSave", WebBaseSave)
	auth("baseGet", WebBaseGet)

	auth("topicCategoryList", WebTopicCategoryList)
	auth("topicCategoryAdd", WebTopicCategoryAdd)
	auth("topicCategoryDel", WebTopicCategoryDel)
	auth("topicCategoryModify", WebTopicCategoryModify)

	auth("zxSave", WebZxSave)
	auth("zxPub", WebZxPub)
	auth("zxGet", WebZxGet)

	auth("topicList", WebTopicList)

	auth("userList", WebUserList)
}

func VerifyAdmin(c *gin.Context) (bool, *Param) {
	param, _ := Parameter(c)
	ect, e := c.Cookie("ecToken")
	ua, eua := c.Cookie("ua")
	if e == nil && eua == nil && len(ect) > 2 && len(ua) > 1 {
		idx := strings.Index(ect, "_")
		if idx > 0 {
			id := ect[0:idx]
			etk := ua + ect[0:idx]
			v, vb := ecTokenMap[etk]
			if !vb || v != ect[idx+1:] {
				qTk, qb, _ := ws.Gpa.QueryString("select Token from site.EcToken where Ua=? and UserId=?", ua, id)
				if !qb {
					return false, param
				} else {
					if qTk == ect[idx+1:] {
						ecTokenMap[etk] = qTk
					} else {
						return false, param
					}
				}
			}

			siteId := param.Int64("siteId", 0)
			if siteId > 0 {
				ext, _, _ := ws.Gpa.QueryInt("select count(*) from site.Site where UserId=? and Id=?", id, siteId)
				if ext == 1 {
					param.SiteId = siteId
					return true, param
				}
			}
		}
	}
	return false, param
}
