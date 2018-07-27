package gktopic

import (
	"strings"

	"fmt"
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebAdd(auth *ws.Web) {
	if auth.ScoreLack() {
		return
	}

	cid := auth.String("CategoryId")
	Id, _ := ws.TopicDao.Add(auth.SiteId, auth.String("Title"), auth.String("Body"),
		auth.UserId, cid, auth.Username, auth.String("SourceType"), auth.String("Source"))
	auth.Out["Id"] = Id
	ws.TopicCategoryDao.UpItemCount(auth.SiteId, cid)
	c, _, _ := ws.TopicDao.CountByUserId(auth.SiteId, auth.UserId)
	sId := fmt.Sprint(Id)

	auth.Out["Score"] = gkuser.UpCount(&ws.UpReq{UserId: auth.UserId, Type: "Topic", Val: c, Fee: ws.GetSoreRule(auth.SiteId).Topic,
		ScoreType: "创建主题", EntityId: sId, SiteId: auth.SiteId,
		ScoreDesc: `创建主题 › <go onclick="vgo('/p/topic/detail,` + sId + `,` + fmt.Sprint(auth.UserId) + `',this)">` + auth.String("Title") + "</go>"})

	ws.KvDao.TopicCount(auth.SiteId, auth.SiteId)

}

//-- 反爬虫，多一个参数UserId
func WebDetail(web *ws.Web) {
	id := web.Int64("id")
	uId := web.String("uId")
	idx := strings.Index(uId, ";")
	if idx > 0 {
		uId = uId[0:idx]
	}
	web.Out["appendList"], _ = ws.AppendDao.List(web.SiteId, id)
	efr, _ := ws.TopicDao.ModifyShowTimes(web.SiteId, id, uId)
	if efr == 1 {
		bx := false
		web.Out["Topic"], bx, _ = ws.TopicDao.FindById(web.SiteId, id, uId)
		if bx {
			web.Out["Reply"], _ = ws.ReplyDao.List(web.SiteId, id)
		}

		if web.Auth {
			web.Out ["thank"], _ = ws.ThankDao.List(web.SiteId, id, web.UserId)
		}
		referer := web.String("referer")
		if len(referer) > 0 && strings.Index(referer, "http") == 0 && len(referer) < 240 {
			ch, chb, _ := ws.RefererDao.Check(web.SiteId, id, referer)
			if chb && ch == 0 {
				ws.RefererDao.Add(web.SiteId, id, referer)
			} else {
				ws.RefererDao.Up(web.SiteId, id, referer)
			}
		}
	} else {
		web.ST(ws.StErrorParameter)
	}
}

func WebList(web *ws.Web) {
	id := web.String("cId")
	pId := web.String("pId")
	s := web.Start()
	if id != "0" {
		web.Out["topicList"], _ = ws.TopicDao.ListBySub(web.SiteId, id, s)
	} else {
		if pId != "0" {
			web.Out["topicList"], _ = ws.TopicDao.ListByParent(web.SiteId, pId, s)
		} else {
			web.Out["topicList"], _ = ws.TopicDao.List(web.SiteId, s)
		}
	}
}

func WebListByUserId(web *ws.Web) {
	uId := web.Int64("userId")
	web.Out["list"], _ = ws.TopicDao.ListByUserId(web.SiteId, uId, web.Start())
}

func WebTopicHot(web *ws.Web) {
	web.Out["today"], _ = ws.TopicDao.HotToday(web.SiteId, )
}

func WebTopicBase(web *ws.Web) {
	web.Out["base"], _, _ = ws.TopicDao.FindBase(web.SiteId, web.Int64("Id"))
}

func WebAppend(auth *ws.Web) {
	if auth.ScoreLack() {
		return
	}
	id := auth.Int64("TopicId")
	at := auth.String("AppendText")
	if id == 0 || len(at) == 0 {
		auth.ST(ws.StErrorParameter)
		return
	}
	base, _, _ := ws.TopicDao.FindBase(auth.SiteId, id)

	if base["UserId"] == fmt.Sprint(auth.UserId) {
		c, cb, _ := ws.AppendDao.Count(auth.SiteId, id)
		if cb && c < 3 {
			appId, _ := ws.AppendDao.Add(auth.SiteId, id, at)
			auth.Out["Id"] = appId
			_, auth.Out["score"], _ = gkuser.ChangeScore(auth.SiteId, "append:"+fmt.Sprint(appId), "创建主题附言",
				"创建了附言 > "+topicLink(base["Id"], base["UserId"], base["Title"]), -20, auth.UserId)
		} else {
			auth.ST(ws.StMax)
		}
	} else {
		auth.ST(ws.StErrorParameter)
		return
	}
}
