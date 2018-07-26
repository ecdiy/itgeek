package ws

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/cihub/seelog"
	"strings"
)



type Web struct {
	param, Out map[string]interface{}
	Ua         string
	Context    *gin.Context
	SiteId     int64
}

func (p *Web) String(n string) string {
	return fmt.Sprint(p.param[ n])
}
func (p *Web) Int64(n string) int64 {
	return p.Int64Default(n, 0)
}
func (p *Web) Int64Default(n string, def int64) int64 {
	i, e := strconv.ParseInt(fmt.Sprint(p.param[ n]), 10, 0)
	if e != nil {
		return def
	}
	return i
}
func (p *Web) Start() int64 {
	return p.StartPageSize(PageSize)
}

func (p *Web) StartPageSize(ps int64) int64 {
	page := p.Int64Default("page", 1)
	if page < 1 {
		return int64(0)
	}
	return (page - 1) * ps
}

//--

type Auth struct {
	Web
	UserId int64
	auth   map[string]interface{}
}

func (p *Auth) Username() string {
	return fmt.Sprint(p.auth["Username"])
}

//--
func (p *Web) ST(st *ST) {
	p.Out["Code"] = st.Code
	p.Out["Msg"] = st.Msg
}

func SiteId(c *gin.Context) int64 {
	if MultiSite == 1 {
		sql := "select SiteId from site.BindDomain where Domain=?"
		host := c.Request.Host
		ix := strings.Index(host, ":")
		if ix > 0 {
			host = host[0:ix]
		}
		idx := strings.Index(host, ".ecdiy.cn")
		if idx > 0 {
			sql = "select Id from site.Site where SubDomain=?"
			host = host[0:idx]
		}
		hId, qb, _ := Gpa.QueryInt64(sql, host)
		if qb {
			return hId
		}
		return 0
	}
	return 0
}

func GetWeb(c *gin.Context) (*Web, error) {
	p, e := c.Get("param")
	if e {
		return p.(*Web), nil
	}
	ua, _ := c.Cookie("ua")
	if ua == "web" {
		ua = "web"
	} else {
		ua = "h5"
	}
	pm := &Web{Ua: ua, Context: c, SiteId: SiteId(c), Out: make(map[string]interface{})}
	row, b := c.GetRawData()
	if b == nil {
		var data map[string]interface{}
		je := json.Unmarshal(row, &data)
		if je != nil {
			seelog.Error("RawData JSON error:", row, je)
			return pm, je
		} else {
			pm.param = data
			c.Set("param", pm)
			return pm, je
		}
	}
	return pm, b
}

func GetAuth(c *gin.Context) (*Auth, error) {
	p, e := c.Get("param")
	if e {
		return p.(*Auth), nil
	}
	ua, _ := c.Cookie("ua")
	if ua == "web" {
		ua = "web"
	} else {
		ua = "h5"
	}
	pm := &Auth{}
	pm.Ua = ua
	pm.Context = c
	pm.SiteId = SiteId(c)
	pm.Out = make(map[string]interface{})
	row, b := c.GetRawData()
	if b == nil {
		var data map[string]interface{}
		je := json.Unmarshal(row, &data)
		if je != nil {
			seelog.Error("RawData JSON error:", row, je)
			return pm, je
		} else {
			pm.param = data
			c.Set("param", pm)
			return pm, je
		}
	}
	return pm, b
}

func WebAuth(Gin *gin.Engine, url string, fun func(wdb *Auth), verify func(c *gin.Context) (bool, int64)) {
	Gin.POST(url, func(c *gin.Context) {
		auth, userId := verify(c)
		if auth {
			param, e := GetAuth(c)
			param.UserId = userId
			if e == nil {
				param.auth = c.Keys
				fun(param)
				c.JSON(200, param.Out)
			}
		} else {
			c.AbortWithStatus(401)
		}
	})
}

func WebPost(Gin *gin.Engine, url string, fun func(wdb *Web)) {
	Gin.POST(url, func(c *gin.Context) {
		param, e := GetWeb(c)
		if e == nil {
			fun(param)
			c.JSON(200, param.Out)
		}
	})
}
