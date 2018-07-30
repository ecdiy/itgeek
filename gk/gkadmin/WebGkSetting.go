package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebBaseGet(web *ws.Web) {
	web.Out["baseInfo"], _, _ = ws.KvDao.Get(web.SiteId, web.String("key"))
}
func WebBaseSave(web *ws.Web) {
	InitCount(web.SiteId)
	key := web.String("key")
	_, b, _ := ws.KvDao.Get(web.SiteId, key)
	if b {
		web.Out["row"], _ = ws.KvDao.Update(web.String("info"), key, web.SiteId)
	} else {
		web.Out["row"], _ = ws.KvDao.Add(web.SiteId, key, web.String("info"))
	}
}
func InitCount(siteId int64) {
	key := []string{"TopicCount", "ReplyCount", "UserCount"}
	for _, v := range key {
		_, b, _ := ws.KvDao.Get(siteId, v)
		if !b {
			ws.KvDao.Add(siteId, v, 0)
		}
	}
}