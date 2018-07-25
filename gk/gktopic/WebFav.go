package gktopic

import (
	"github.com/ecdiy/itgeek/gk/gkuser"
	"github.com/ecdiy/itgeek/gk/ws"
)

func WebFavList(UserId int64, param *ws.Param, res map[string]interface{}) {
	page := param.Int64("page", 1)
	if page >= 1 {
		page = (page - 1) * 20
	}
	res["FavList"], _ = ws.FavDao.List(param.SiteId, UserId, page)
}

func WebFav(UserId int64, param *ws.Param, res map[string]interface{}) {
	id := param.String("id")
	fc, _, _ := ws.FavDao.Exist(param.SiteId, UserId, id)
	doType := 1
	if fc == 1 {
		ws.FavDao.Del(param.SiteId, UserId, id)
		doType = 0
	} else {
		ws.FavDao.Save(param.SiteId, UserId, id)
	}
	tfc, _, _ := ws.FavDao.Count(param.SiteId, UserId)

	gkuser.UpCount(&ws.UpReq{UserId: UserId, Type: "Fav", Val: tfc, Fee: int64(0), SiteId: param.SiteId,})
	res["TopicFavCount"] = tfc
	res["Fav"] = doType
}

func WebFavStatus(UserId int64, param *ws.Param, res map[string]interface{}) {
	res["Fav"], _, _ = ws.FavDao.Exist(param.SiteId, UserId, param.Int64("id", 0))
}
