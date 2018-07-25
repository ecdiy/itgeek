package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebBaseGet(param *Param, res map[string]interface{}) {
	res["baseInfo"], _, _ = ws.KvDao.Get(param.SiteId, param.String("key"))
}
func WebBaseSave(param *Param, res map[string]interface{}) {
	InitCount(param.SiteId)
	key := param.String("key")
	_, b, _ := ws.KvDao.Get(param.SiteId, key)
	if b {
		res["row"], _ = ws.KvDao.Update(param.String("info"), key, param.SiteId)
	} else {
		res["row"], _ = ws.KvDao.Add(param.SiteId, key, param.String("info"))
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
