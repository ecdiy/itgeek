package gkuser

import (
	"strings"
	"github.com/cihub/seelog"
	"strconv"

	"github.com/ecdiy/itgeek/gk/ws"
)

func doVerify(token, ua string, siteId int64) (bool, int64, string) {
	if len(token) > 1 {
		idx := strings.Index(token, "_")
		if idx > 0 {
			id := token[0:idx]
			v, b := tokenMap [ua+"_"+id]
			uid, _ := strconv.ParseInt(id, 10, 0)
			if b {
				if v["Token"] == token[idx+1:] {
					ws.UserDao.DauAdd(siteId, uid)
					return true, uid, v["Username"]
				}
			} else {
				tk, b, ee := ws.TokenDao.Find(siteId, uid, ua)
				if ee == nil && b && tk == token[idx+1:] {
					vf, _, _ := ws.UserDao.BaseInfo(siteId, id)
					vf["Token"] = tk
					tokenMap[ua+"_"+id] = vf
					ws.UserDao.DauAdd(siteId, uid)
					return true, uid, vf["Username"]
				}
			}
		} else {
			seelog.Error("error token.", token)
		}
	}
	return false, 0, ""
}

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
		_, sc, _ := ChangeScore(in.SiteId, in.EntityId, in.ScoreType, in.ScoreDesc, in.Fee, in.UserId)
		return int64(sc)
	}
	return 0
}

func Msg(in *ws.MsgReq) (*ws.NilResp, error) {
	ws.MsgDao.Add(in.SiteId, in.Title, in.Body, in.EntityId, in.MsgType, in.FromUserId, in.UserId, in.GroupId)
	ws.UserDao.UpMsg(in.UserId, in.UserId, in.SiteId)
	return &ws.NilResp{}, nil
}
