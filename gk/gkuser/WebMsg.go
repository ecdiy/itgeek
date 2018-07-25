package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebMsgList(userId int64, param *ws.Param, res map[string]interface{}) {
	res["totalAll"], _, _ = ws.MsgDao.Count(param.SiteId, userId)
	res["totalUnread"], _, _ = ws.MsgDao.CountUnread(param.SiteId, userId)
	//res["totalReply"], _, _ = MsgDao.CountByType(userId, "主题回复")

	mType := param.String("mType")
	if mType == "all" {
		res["msgList"], _ = ws.MsgDao.List(param.SiteId, userId, param.Start(20))
		res["total"] = res["totalAll"]
	}
	if mType == "unRead" {
		res["msgList"], _ = ws.MsgDao.ListUnread(param.SiteId, userId, param.Start(20))
		res["total"] = res["totalUnread"]
	}
	//if mType == "reply" {
	//	res["msgList"], _ = MsgDao.ListByType(userId, "主题回复", param.Start(20))
	//	res["total"] = res["totalAll"]
	//}
}

func WebMsgDel(userId int64, param *ws.Param, res map[string]interface{}) {

	res["row"], _ = ws.MsgDao.Del(param.SiteId, param.Int64("id", 0), userId)
	ws.UserDao.UpMsg(userId, userId, param.SiteId)
	res["gkUser"], _, _ = ws.UserDao.BaseInfo(param.SiteId, userId)
	res["totalAll"], _, _ = ws.MsgDao.Count(param.SiteId, userId)
	res["totalUnread"], _, _ = ws.MsgDao.CountUnread(param.SiteId, userId)
}

func WebMsgRead(userId int64, param *ws.Param, res map[string]interface{}) {
	id := param.Int64("id", 0)
	gId, gb, _ := ws.MsgDao.FindGroupId(param.SiteId, userId, id)
	if gb {
		if gId > 0 {
			res["row"], _ = ws.MsgDao.ReadGroup(param.SiteId, userId, gId)
		} else {
			res["row"], _ = ws.MsgDao.Read(param.SiteId, userId, id)
		}
	}
	ws.UserDao.UpMsg(userId, userId, param.SiteId)
	res["gkUser"], _, _ = ws.UserDao.BaseInfo(param.SiteId, userId)
}
