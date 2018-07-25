package ws

import (
	"github.com/cihub/seelog"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type Result struct {
	Status *ST
	Result string
	Param  map[string]string
}

func (c *Result) ParamToJson() string {
	js1, jse := json.Marshal(c.Param)
	if jse != nil {
		seelog.Error("JSON Format error.", jse)
	}
	return string(js1)
}

//--
func (c *Result) Put(key, val string) *Result {
	if c.Param == nil {
		c.Param = make(map[string]string)
	}
	c.Param[key] = val
	return c
}
func (c *Result) PutAll(m map[string]string) *Result {
	if c.Param == nil {
		c.Param = make(map[string]string)
	}
	for k, v := range m {
		c.Param[k] = v
	}
	return c
}

func (c *Result) OutJSON(gin *gin.Context) {
	gin.Header("Content-Type", "application/json; charset=utf-8")
	gin.JSON(http.StatusOK, c)
}

func (c *Result) OutJsonP(gin *gin.Context) {
	gin.Header("Content-Type", "application/javascript; charset=utf-8")
	callback, b := gin.GetQuery("callback")
	if b {
		gin.HTML(http.StatusOK, "application/javascript; charset=utf-8", callback+"("+fmt.Sprint(c)+")")
	} else {
		gin.Header("Content-Type", "application/json; charset=utf-8")
		gin.JSON(http.StatusOK, c)
	}
}
