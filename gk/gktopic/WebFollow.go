package gktopic

import (
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebFollow(auth *ws.Auth) {
	id := auth.Int64("userId")
	fc, _, _ := ws.FollowDao.Exist(auth.SiteId, id, auth.UserId)
	doType := 1
	if fc == 1 {
		ws.FollowDao.UnFollow(auth.SiteId, id, auth.UserId)
		doType = 0
	} else {
		ws.FollowDao.Follow(auth.SiteId, id, auth.UserId)
	}
	tfc, _, _ := ws.FollowDao.Count(auth.SiteId, auth.UserId)

	gkuser.UpCount(&ws.UpReq{
		UserId: auth.UserId, Type: "Follow", Val: tfc, Fee: int64(0),
	})

	auth.Out["followCount"] = tfc
	auth.Out["followStatus"] = doType
}

func WebFollowList(auth *ws.Auth) {

	auth.Out["FavList"], _ = ws.FollowDao.TopicList(auth.SiteId, auth.UserId, auth.Start())
}
func WebFollowStatus(auth *ws.Auth) {
	auth.Out["followStatus"], _, _ = ws.FollowDao.Exist(auth.SiteId, auth.Int64("id"), auth.UserId)
}
