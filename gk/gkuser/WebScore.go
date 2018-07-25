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

func WebScoreLoginAward(userId int64, param *ws.Param, res map[string]interface{}) {
	res["LoginAward"], _, _ = ws.UserDao.LoginAward(param.SiteId, userId)
}
func WebScoreLoginAwardDo(userId int64, param *ws.Param, res map[string]interface{}) {
	val := GenerateRangeNum(5, 50)
	t := time.Now().Format("2006-01-02")
	id, score, _ := ChangeScore(param.SiteId, t, "每日登录奖励", t+"的每日登录奖励", val, userId)
	if id > 0 {
		ws.UserDao.LoginAwardDo(param.SiteId, userId)
	}
	res["Id"] = id
	res["Score"] = score
}

func WebScoreLogList(userId int64, param *ws.Param, res map[string]interface{}) {
	page := param.Int64("page", 1)
	page = (page - 1) * 20
	res["total"], _, _ = ws.ScoreLog.Count(param.SiteId, userId)
	res["list"], _ = ws.ScoreLog.List(param.SiteId, userId, page)
}
