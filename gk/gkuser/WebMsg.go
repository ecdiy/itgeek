package gkuser

import (
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebMsgList(auth *ws.Auth) {
	auth.Out["totalAll"], _, _ = ws.MsgDao.Count(auth.SiteId, auth.UserId)
	auth.Out["totalUnread"], _, _ = ws.MsgDao.CountUnread(auth.SiteId, auth.UserId)
	//auth.Out["totalReply"], _, _ = MsgDao.CountByType(auth.UserId, "主题回复")

	mType := auth.String("mType")
	if mType == "all" {
		auth.Out["msgList"], _ = ws.MsgDao.List(auth.SiteId, auth.UserId, auth.Start())
		auth.Out["total"] = auth.Out["totalAll"]
	}
	if mType == "unRead" {
		auth.Out["msgList"], _ = ws.MsgDao.ListUnread(auth.SiteId, auth.UserId, auth.Start())
		auth.Out["total"] = auth.Out["totalUnread"]
	}
	//if mType == "reply" {
	//	auth.Out["msgList"], _ = MsgDao.ListByType(auth.UserId, "主题回复", auth.Start(20))
	//	auth.Out["total"] = auth.Out["totalAll"]
	//}
}

func WebMsgDel(auth *ws.Auth) {

	auth.Out["row"], _ = ws.MsgDao.Del(auth.SiteId, auth.Int64("id"), auth.UserId)
	ws.UserDao.UpMsg(auth.UserId, auth.UserId, auth.SiteId)
	auth.Out["gkUser"], _, _ = ws.UserDao.BaseInfo(auth.SiteId, auth.UserId)
	auth.Out["totalAll"], _, _ = ws.MsgDao.Count(auth.SiteId, auth.UserId)
	auth.Out["totalUnread"], _, _ = ws.MsgDao.CountUnread(auth.SiteId, auth.UserId)
}

func WebMsgRead(auth *ws.Auth) {
	id := auth.Int64("id")
	gId, gb, _ := ws.MsgDao.FindGroupId(auth.SiteId, auth.UserId, id)
	if gb {
		if gId > 0 {
			auth.Out["row"], _ = ws.MsgDao.ReadGroup(auth.SiteId, auth.UserId, gId)
		} else {
			auth.Out["row"], _ = ws.MsgDao.Read(auth.SiteId, auth.UserId, id)
		}
	}
	ws.UserDao.UpMsg(auth.UserId, auth.UserId, auth.SiteId)
	auth.Out["gkUser"], _, _ = ws.UserDao.BaseInfo(auth.SiteId, auth.UserId)
}
