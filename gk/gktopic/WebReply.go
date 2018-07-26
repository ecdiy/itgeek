package gktopic

import (
	"fmt"
	"strconv"
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebReply(auth *ws.Auth) {
	if auth.ScoreLack() {
		return
	}

	tId := auth.Int64("TopicId")
	id, _ := ws.ReplyDao.Insert(auth.SiteId, auth.UserId, tId, auth.Username(), auth.String("Reply"))
	auth.Out["Id"] = id
	cr, _, _ := ws.ReplyDao.Count(auth.SiteId, tId)
	ws.TopicDao.UpdateReply(auth.Username(), cr, auth.UserId, tId)
	auth.Out["Count"] = cr
	auth.Out["list"], _ = ws.ReplyDao.List(auth.SiteId, tId)
	c, _, _ := ws.ReplyDao.CountByUserId(auth.SiteId, auth.UserId)
	tp, _, _ := ws.TopicDao.Find(tId)
	createId, _ := strconv.ParseInt(tp["UserId"], 10, 0)

	entityId := "Reply:" + fmt.Sprint(id)

	topicLink := ` <a onclick="vgo('/p/topic/detail,` + fmt.Sprint(tId) + `,` + tp["UserId"] + `',this)">` + tp["Title"] + `</a>`
	memberLink := memberLink(auth.Username())

	sr := ws.GetSoreRule(auth.SiteId)

	gkuser.UpCount(&ws.UpReq{UserId: auth.UserId, Type: "Reply", Val: c, Fee: sr.Reply, SiteId: auth.SiteId,
		ScoreType: "创建回复", EntityId: fmt.Sprint(id),
		ScoreDesc: `创建复 ›  ` + topicLink})

	if auth.UserId != createId {

		gkuser.UpCount(&ws.UpReq{UserId: createId, Fee: sr.GetReply, EntityId: fmt.Sprint(id), SiteId: auth.SiteId,
			ScoreType: "主题回复收益",
			ScoreDesc: `收到 ` + memberLink + ` 的回复 › ` + topicLink})

		gkuser.Msg(&ws.MsgReq{EntityId: entityId, MsgType: "主题回复", GroupId: tId,
			FromUserId: auth.UserId, UserId: createId, Body: auth.String("Reply"), FromUsername: auth.Username(),
			SiteId: auth.SiteId, Title: "在 " + topicLink + ` 里回复了你`,
		})

	}
	auth.Out["ScoreLack"], _, _ = ws.UserDao.Score(auth.SiteId, auth.UserId)
}


func WebTopicReplyThank(auth *ws.Auth) {
	if auth.ScoreLack() {
		return
	}
	rId := auth.Int64("ReplyId")
	replyInfo, fd, fdErr := ws.ReplyDao.Get(auth.SiteId, rId)
	if !fd || fdErr != nil {
		return
	}
	toId, _ := strconv.ParseInt(fmt.Sprint(replyInfo["UserId"]), 10, 0)
	e, _, _ := ws.ThankDao.Exist(auth.SiteId, rId)
	topicId, _ := strconv.ParseInt(replyInfo["TopicId"], 10, 0)
	if e == 0 && toId != auth.UserId {
		thId, _ := ws.ThankDao.Add(auth.SiteId, auth.UserId, rId, replyInfo["TopicId"])
		tl := topicLink(replyInfo["TopicId"], replyInfo["Author"], replyInfo["Title"])
		sr := ws.GetSoreRule(auth.SiteId)

		_, auth.Out["score"], _ = gkuser.ChangeScore(auth.SiteId, "thank:"+strconv.FormatInt(thId, 10), "发送谢意",
			"感谢 "+memberLink(replyInfo["Username"])+" 的回复 > "+tl, sr.Thank, auth.UserId)

		entityId := "thank:" + strconv.FormatInt(thId, 10)

		gkuser.ChangeScore(auth.SiteId, entityId, "收到谢意",
			"感谢 "+memberLink(auth.Username())+" 的回复", sr.Thank*(-1), toId)

		gkuser.Msg(&ws.MsgReq{EntityId: entityId, MsgType: "收到谢意", GroupId: topicId,
			FromUserId: auth.UserId, UserId: toId, Body: auth.String("Reply"), FromUsername: auth.Username(),
			SiteId: auth.SiteId, Title: memberLink(auth.Username()) + " 感谢你在 " + tl + ` 的回复`,})
	}
}
