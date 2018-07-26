package gktopic

import (
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebFavList(auth *ws.Auth) {

	auth.Out["FavList"], _ = ws.FavDao.List(auth.SiteId, auth.UserId, auth.Start())
}

func WebFav(auth *ws.Auth) {
	id := auth.String("id")
	fc, _, _ := ws.FavDao.Exist(auth.SiteId, auth.UserId, id)
	doType := 1
	if fc == 1 {
		ws.FavDao.Del(auth.SiteId, auth.UserId, id)
		doType = 0
	} else {
		ws.FavDao.Save(auth.SiteId, auth.UserId, id)
	}
	tfc, _, _ := ws.FavDao.Count(auth.SiteId, auth.UserId)

	gkuser.UpCount(&ws.UpReq{UserId: auth.UserId, Type: "Fav", Val: tfc, Fee: int64(0), SiteId: auth.SiteId,})
	auth.Out["TopicFavCount"] = tfc
	auth.Out["Fav"] = doType
}

func WebFavStatus(auth *ws.Auth) {
	auth.Out["Fav"], _, _ = ws.FavDao.Exist(auth.SiteId, auth.UserId, auth.Int64("id"))
}
