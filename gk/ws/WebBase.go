package ws

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"encoding/json"
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

//--

func (p *Web) ScoreLack() bool { //检查积分
	sc, _, _ := UserDao.Score(p.SiteId, p.UserId)
	if sc < 0 {
		p.ST(StScoreLack)
		return true
	}
	return false
}

//--
func (p *Web) Result(result ... interface{}) {
	if result != nil {
		if len(result) == 1 {
			p.Out["result"] = result[0]
		} else {
			p.Out["result"] = result
		}
	}
}

func (p *Web) ST(st *ST, result ... interface{}) {
	p.Out["code"] = st.Code
	p.Out["msg"] = st.Msg
	p.Result(result ...)
}
