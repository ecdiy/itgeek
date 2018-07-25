package gktopic

import (
	"strings"
	"utils/base"

	"fmt"
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebAdd(UserId int64, param *ws.Param, res map[string]interface{}) {
	sc, _, _ := ws.UserDao.Score(param.SiteId, UserId)
	if sc < 0 {
		res["error"] = "积分不够，不能创建"
		return
	}

	cid := param.String("CategoryId")
	Id, _ := ws.TopicDao.Add(param.SiteId, param.String("Title"), param.String("Body"),
		UserId, cid, param.Username(), param.String("SourceType"), param.String("Source"))
	res["Id"] = Id
	ws.TopicCategoryDao.UpItemCount(param.SiteId, cid)
	c, _, _ := ws.TopicDao.CountByUserId(param.SiteId, UserId)
	sId := fmt.Sprint(Id)

	res["Score"] = gkuser.UpCount(&ws.UpReq{UserId: UserId, Type: "Topic", Val: c, Fee: ws.GetSoreRule(param.SiteId).Topic,
		ScoreType: "创建主题", EntityId: sId, SiteId: param.SiteId,
		ScoreDesc: `创建主题 › <go onclick="vgo('/p/topic/detail,` + sId + `,` + fmt.Sprint(UserId) + `',this)">` + param.String("Title") + "</go>"})

	ws.KvDao.TopicCount(param.SiteId, param.SiteId)

}

//-- 反爬虫，多一个参数UserId
func WebDetail(param *ws.Param, res map[string]interface{}) {
	id := param.Int64("id", 0)
	uId := param.String("uId")
	idx := strings.Index(uId, ";")
	if idx > 0 {
		uId = uId[0:idx]
	}
	//seelog.Info(id, ",", uId)
	efr, _ := ws.TopicDao.ModifyShowTimes(param.SiteId, id, uId)
	if efr == 1 {
		bx := false
		res["Topic"], bx, _ = ws.TopicDao.FindById(param.SiteId, id, uId)
		if bx {
			res["Reply"], _ = ws.ReplyDao.List(param.SiteId, id)
		}
		b, bu := gkuser.Verify(param.Context)
		if b {
			res["thank"], _ = ws.ThankDao.List(param.SiteId, id, bu)
		}
		referer := param.String("referer")
		if len(referer) > 0 {
			ch, chb, _ := ws.RefererDao.Check(param.SiteId, id, referer)
			if chb && ch == 0 {
				ws.RefererDao.Add(param.SiteId, id, referer)
			} else {
				ws.RefererDao.Up(param.SiteId, id, referer)
			}
		}
	} else {
		base.StErrorParameter.To(res)
	}
}

func WebList(param *ws.Param, res map[string]interface{}) {
	id := param.String("cId")
	pId := param.String("pId")
	page := param.Int64("page", 1)
	page = (page - 1) * 20
	if id != "0" {
		res["topicList"], _ = ws.TopicDao.ListBySub(param.SiteId, id, page)
	} else {
		if pId != "0" {
			res["topicList"], _ = ws.TopicDao.ListByParent(param.SiteId, pId, page)
		} else {
			res["topicList"], _ = ws.TopicDao.List(param.SiteId, page)
		}
	}
}

func WebListByUserId(param *ws.Param, res map[string]interface{}) {
	username := param.Int64("userId", 0)
	page := param.Int64("page", 1)
	page = (page - 1) * 20
	res["list"], _ = ws.TopicDao.ListByUserId(param.SiteId, username, page)
}

func WebTopicHot(param *ws.Param, res map[string]interface{}) {
	res["today"], _ = ws.TopicDao.HotToday(param.SiteId, )
	//res["all"], _ = TopicDao.HotAll()
}
