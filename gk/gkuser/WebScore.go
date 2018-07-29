package gkuser

import (
	"math/rand"
	"time"
	"github.com/ecdiy/itgeek/gk/ws"
)

func ChangeScore(siteId int64, entityId, scoreType, scoreDesc string, fee int64, userId int64) (int64, int64, error) {
	sc, _, _ := ws.UserDao.Score(siteId, userId)
	score := fee + sc
	id, e := ws.ScoreLog.Add(siteId, score, scoreType, scoreDesc, entityId, fee, userId)
	if id > 0 {
		ws.UserDao.UpScore(score, userId, siteId)
	}
	return id, score, e
}

func GenerateRangeNum(min, max int) int64 {
	rand.Seed(time.Now().Unix())
	return int64(rand.Intn(max-min) + min)
}

func WebScoreLoginAwardStatus(auth *ws.Web) {
	//LoginAward,LoginDay,AwardDate
	res, rb, e := ws.UserDao.LoginAward(auth.SiteId, auth.UserId)
	if e == nil && rb {
		auth.Out["LoginAward"] = res["LoginAward"]
		auth.Out["today"] = time.Now().Format("2006-01-02")
		auth.Out["loginDay"] = res["LoginDay"]
		//if auth.Out["LoginAward"] == "1" {
		//
		//}
	}
}
func WebScoreLoginAwardDo(auth *ws.Web) {
	val := GenerateRangeNum(5, 50)
	t := time.Now().Format("2006-01-02")
	id, score, _ := ChangeScore(auth.SiteId, t, "每日登录奖励", t+"的每日登录奖励", val, auth.UserId)
	if id > 0 {
		ws.UserDao.LoginAwardDo(auth.SiteId, auth.UserId)
	}
	auth.Out["Id"] = id
	auth.Out["Score"] = score
	WebScoreLoginAwardStatus(auth)
}

func WebScoreLogList(auth *ws.Web) {
	auth.Out["total"], _, _ = ws.ScoreLog.Count(auth.SiteId, auth.UserId)
	auth.Out["list"], _ = ws.ScoreLog.List(auth.SiteId, auth.UserId, auth.Start())
}
