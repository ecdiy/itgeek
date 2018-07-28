package ws

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"encoding/json"
	"strings"
)

type Web struct {
	param, Out map[string]interface{}
	Ua         string
	Context    *gin.Context
	SiteId     int64
	//---
	Auth     bool
	UserId   int64
	auth     map[string]interface{}
	Username string
}

func (p *Web) initParam() {
	row, b := p.Context.GetRawData()
	if b == nil {
		var data map[string]interface{}
		je := json.Unmarshal(row, &data)
		if je == nil {
			p.param = data
		}
	}
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
func (p *Web) Result(result ... interface{}) {
	if result != nil {
		p.Out["result"] = result
	}
}
func (p *Web) ST(st *ST, result ... interface{}) {
	p.Out["code"] = st.Code
	p.Out["msg"] = st.Msg
	p.Result(result ...)
}
func (p *Web) ScoreLack() bool { //检查积分
	sc, _, _ := UserDao.Score(p.SiteId, p.UserId)
	if sc < 0 {
		p.ST(StScoreLack)
		return true
	}
	return false
}

//--

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

func WebAuth(url string, fun func(wdb *Web)) {
	WebGin.POST(url, func(c *gin.Context) {
		web := Verify(c)
		if web.Auth {
			web.initParam()
			fun(web)
			c.JSON(200, web.Out)
		} else {
			c.AbortWithStatus(401)
		}
	})
}

func WebPost(url string, fun func(wdb *Web)) {
	WebGin.POST(url, func(c *gin.Context) {
		web := Verify(c)
		web.initParam()
		fun(web)
		c.JSON(200, web.Out)
	})
}

func GetUa(ctx *gin.Context) string {
	ua := ctx.Request.UserAgent()
	if len(ua) == 0 {
		return "web"
	}
	if UaH5.MatchString(ua) {
		return "h5"
	}
	if UaSeo.MatchString(ua) {
		return "seo"
	}
	return "web"
}
