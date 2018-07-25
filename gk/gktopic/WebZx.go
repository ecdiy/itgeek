package gktopic

import "github.com/ecdiy/itgeek/gk/ws"

func WebZxPage(param *ws.Param, res map[string]interface{}) {
	res["result"], _, _ = ws.ZxDao.Get(param.SiteId, param.Ua, param.String("PageKey"))
}
