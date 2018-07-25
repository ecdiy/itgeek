package gktopic

import (
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebFollow(userId int64, param *ws.Param, res map[string]interface{}) {
	id := param.Int64("userId", 0)
	fc, _, _ := ws.FollowDao.Exist(param.SiteId, id, userId)
	doType := 1
	if fc == 1 {
		ws.FollowDao.UnFollow(param.SiteId, id, userId)
		doType = 0
	} else {
		ws.FollowDao.Follow(param.SiteId, id, userId)
	}
	tfc, _, _ := ws.FollowDao.Count(param.SiteId, userId)

	gkuser.UpCount(&ws.UpReq{
		UserId: userId, Type: "Follow", Val: tfc, Fee: int64(0),
	})

	res["followCount"] = tfc
	res["followStatus"] = doType
}

func WebFollowList(userId int64, param *ws.Param, res map[string]interface{}) {
	page := param.Int64("page", 1)
	if page >= 1 {
		page = (page - 1) * 20
	}
	res["FavList"], _ = ws.FollowDao.TopicList(param.SiteId, userId, page)
}
func WebFollowStatus(userId int64, param *ws.Param, res map[string]interface{}) {
	res["followStatus"], _, _ = ws.FollowDao.Exist(param.SiteId, param.Int64("id", 0), userId)
}
