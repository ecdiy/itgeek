package gkadmin

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

var ecTokenMap = make(map[string]string)

func InitWeb() {

	ws.WebGin.POST("/api/gk-admin/siteUpload", WebUpload)

	auth := func(url string, fun func(auth *ws.Web)) {
		ws.WebAuthAdmin("/api/gk-admin/"+url, fun)
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
