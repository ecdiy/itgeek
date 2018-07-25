package gktopic

import (
	"fmt"
	"strconv"
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebReply(UserId int64, param *ws.Param, res map[string]interface{}) {
	sc, _, _ := ws.UserDao.Score(param.SiteId, UserId)
	if sc < 0 {
		res["error"] = "积分不够，不能回复"
		return
	}

	tId := param.Int64("TopicId", 0)
	id, _ := ws.ReplyDao.Insert(param.SiteId, UserId, tId, param.Username(), param.String("Reply"))
	res["Id"] = id
	cr, _, _ := ws.ReplyDao.Count(param.SiteId, tId)
	ws.TopicDao.UpdateReply(param.Username(), cr, UserId, tId)
	res["Count"] = cr
	res["list"], _ = ws.ReplyDao.List(param.SiteId, tId)
	c, _, _ := ws.ReplyDao.CountByUserId(param.SiteId, UserId)
	tp, _, _ := ws.TopicDao.Find(tId)
	createId, _ := strconv.ParseInt(tp["UserId"], 10, 0)

	entityId := "Reply:" + fmt.Sprint(id)

	topicLink := ` <a onclick="vgo('/p/topic/detail,` + fmt.Sprint(tId) + `,` + tp["UserId"] + `',this)">` + tp["Title"] + `</a>`
	memberLink := memberLink(param.Username())

	sr := ws.GetSoreRule(param.SiteId)

	gkuser.UpCount(&ws.UpReq{UserId: UserId, Type: "Reply", Val: c, Fee: sr.Reply, SiteId: param.SiteId,
		ScoreType: "创建回复", EntityId: fmt.Sprint(id),
		ScoreDesc: `创建复 ›  ` + topicLink})

	if UserId != createId {

		gkuser.UpCount(&ws.UpReq{UserId: createId, Fee: sr.GetReply, EntityId: fmt.Sprint(id), SiteId: param.SiteId,
			ScoreType: "主题回复收益",
			ScoreDesc: `收到 ` + memberLink + ` 的回复 › ` + topicLink})

		gkuser.Msg(&ws.MsgReq{EntityId: entityId, MsgType: "主题回复", GroupId: tId,
			FromUserId: UserId, UserId: createId, Body: param.String("Reply"), FromUsername: param.Username(),
			SiteId: param.SiteId, Title: "在 " + topicLink + ` 里回复了你`,
		})

	}
	res["Score"], _, _ = ws.UserDao.Score(param.SiteId, UserId)
}

func topicLink(topicId, userId, title string) string {
	return ` <a onclick="vgo('/p/topic/detail,` + topicId + `,` + userId + `',this)">` + title + `</a> `
}
func memberLink(un string) string {
	return ` <a onclick="vgo('/p/member/` + un + `')">` + un + `</a> `
}
func WebTopicReplyThank(UserId int64, param *ws.Param, res map[string]interface{}) {
	rId := param.Int64("ReplyId", 0)
	replyInfo, fd, fdErr := ws.ReplyDao.Get(param.SiteId, rId)
	if !fd || fdErr != nil {
		return
	}
	toId, _ := strconv.ParseInt(fmt.Sprint(replyInfo["UserId"]), 10, 0)
	e, _, _ := ws.ThankDao.Exist(param.SiteId, rId)
	topicId, _ := strconv.ParseInt(replyInfo["TopicId"], 10, 0)
	if e == 0 && toId != UserId {
		thId, _ := ws.ThankDao.Add(param.SiteId, UserId, rId, replyInfo["TopicId"])
		tl := topicLink(replyInfo["TopicId"], replyInfo["Author"], replyInfo["Title"])
		sr := ws.GetSoreRule(param.SiteId)

		_, res["score"], _ = gkuser.ChangeScore(param.SiteId, "thank:"+strconv.FormatInt(thId, 10), "发送谢意",
			"感谢 "+memberLink(replyInfo["Username"])+" 的回复 > "+tl, sr.Thank, UserId)

		entityId := "thank:" + strconv.FormatInt(thId, 10)

		gkuser.ChangeScore(param.SiteId, entityId, "收到谢意",
			"感谢 "+memberLink(param.Username())+" 的回复", sr.Thank*(-1), toId)

		gkuser.Msg(&ws.MsgReq{EntityId: entityId, MsgType: "收到谢意", GroupId: topicId,
			FromUserId: UserId, UserId: toId, Body: param.String("Reply"), FromUsername: param.Username(),
			SiteId: param.SiteId, Title: memberLink(param.Username()) + " 感谢你在 " + tl + ` 的回复`,})
	}
}
