package gkuser

import (
	"github.com/cihub/seelog"
	"github.com/ecdiy/itgeek/gk/ws"
)

func UpCount(in *ws.UpReq) (int64) {
	if in.Fee != 0 && len(in.EntityId) == 0 {
		seelog.Error("参数错误(Fee != 0 && len(in.EntityId) > 0):", in.Fee, in.EntityId)
		return 0
	}
	if in.Type == "Follow" {
		ws.UserDao.UpFollow(in.Val, in.UserId, in.SiteId)
	}
	if in.Type == "Fav" {
		ws.UserDao.UpFav(in.Val, in.UserId, in.SiteId)
	}
	if in.Type == "Reply" {
		ws.UserDao.UpReply(in.Val, in.UserId, in.SiteId)
		ws.KvDao.ReplyCount(in.SiteId, in.SiteId)
	}
	if in.Type == "Topic" {
		ws.UserDao.UpTopic(in.Val, in.UserId, in.SiteId)
		ws.KvDao.TopicCount(in.SiteId, in.SiteId)
	}
	if in.Fee != 0 {
		_, sc, _, _ := ChangeScore(in.SiteId, in.EntityId, in.ScoreType, in.ScoreDesc, in.Fee, in.UserId)
		return int64(sc)
	}
	return 0
}

func Msg(in *ws.MsgReq) {
	ws.MsgDao.Add(in.SiteId, in.Title, in.Body, in.EntityId, in.MsgType, in.FromUserId, in.UserId, in.GroupId)
	ws.UserDao.UpMsg(in.UserId, in.UserId, in.SiteId)
}
