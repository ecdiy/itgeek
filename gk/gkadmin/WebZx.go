package gkadmin

import "github.com/ecdiy/itgeek/gk/ws"

func WebZxSave(param *Param, res map[string]interface{}) {
	_, ext, _ := ws.ZxDao.GetTest(param.SiteId, param.String("PageUa"), param.String("PageKey"))
	if ext {
		res["row"], _ = ws.ZxDao.UpJson(param.String("Json"), param.String("PageUa"), param.String("PageKey"), param.SiteId)
	} else {
		res["row"], _ = ws.ZxDao.Add(param.SiteId, param.String("PageKey"), param.String("PageUa"), param.String("Json"))
	}
}
func WebZxPub(param *Param, res map[string]interface{}) {
	res["row"], _ = ws.ZxDao.Pub(param.String("PageUa"), param.String("PageKey"), param.SiteId)
}
func WebZxGet(param *Param, res map[string]interface{}) {
	res["zxJson"], _, _ = ws.ZxDao.GetTest(param.SiteId, param.String("PageUa"), param.String("PageKey"))
}
