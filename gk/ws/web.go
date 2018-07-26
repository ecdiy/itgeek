package ws



//type Param struct {
//	auth, param map[string]interface{}
//	Ua, uri     string
//	Context     *gin.Context
//	SiteId      int64
//}
//
//func (p *Param) v(n string) interface{} {
//	vx, vb := p.param[p.uri+n]
//	if vb {
//		return vx
//	}
//	return p.param[ n]
//}
//
//func (p *Param) String(n string) string {
//	return fmt.Sprint(p.v(n))
//}
//func (p *Param) Username() string {
//	return fmt.Sprint(p.auth["Username"])
//}
//func (p *Param) Int64(n string, def int64) int64 {
//	i, e := strconv.ParseInt(fmt.Sprint(p.v(n)), 10, 0)
//	if e != nil {
//		return def
//	}
//	return i
//}
//
//func (p *Param) Start(ps int64) int64 {
//	page := p.Int64("page", 1)
//	if page < 1 {
//		return int64(0)
//	}
//	return (page - 1) * ps
//}


//func Parameter(c *gin.Context) (*Param, error) {
//	p, e := c.Get("param")
//	if e {
//		return p.(*Param), nil
//	}
//	ua, _ := c.Cookie("ua")
//	if ua == "web" {
//		ua = "web"
//	} else {
//		ua = "h5"
//	}
//	pm := &Param{Ua: ua, Context: c, SiteId: SiteId(c)}
//	row, b := c.GetRawData()
//	if b == nil {
//		var data map[string]interface{}
//		je := json.Unmarshal(row, &data)
//		if je != nil {
//			seelog.Error("RawData JSON error:", row, je)
//			return pm, je
//		} else {
//			pm.param = data
//			c.Set("param", pm)
//			return pm, je
//		}
//	}
//	return pm, b
//}

//func HttpPost(Gin *gin.Engine, url string, fun func(param *Param, res map[string]interface{})) {
//	funMap[url[1:]] = fun
//	Gin.POST(url, func(c *gin.Context) {
//		param, e := Parameter(c)
//		if e == nil {
//			res := make(map[string]interface{})
//			fun(param, res)
//			c.JSON(200, res)
//		} else {
//			c.AbortWithStatus(404)
//		}
//	})
//}
//
//func HttpAuth(Gin *gin.Engine, url string, fun func(userId int64, param *Param, res map[string]interface{}), verify func(c *gin.Context) (bool, int64)) {
//	authMap[url[1:]] = fun
//	Gin.POST(url, func(c *gin.Context) {
//		auth, userId := verify(c)
//		if auth {
//			param, e := Parameter(c)
//			if e == nil {
//				param.auth = c.Keys
//				res := make(map[string]interface{})
//				fun(userId, param, res)
//				c.JSON(200, res)
//			}
//		} else {
//			c.AbortWithStatus(401)
//		}
//	})
//}
