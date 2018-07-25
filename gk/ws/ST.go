package ws

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	/**
	NE  : NOT EXIST
	Err : ERROR
	Dis ：disable
	 */
	OK = st(0, "")
	//---0~9 全局系统错误
	StErrorParameter  = st(1, "参数错误")
	StErrorToken      = st(2, "用户未登录，Token失效")
	StNotOrderService = st(3, "没有订购服务")
	StErrorRpc        = st(4, "远程RPC调用错误")
	StExist           = st(5, "已存在")
	StNotExist        = st(6, "不存在")
	StErrorDb         = st(7, "数据逻辑错误")
	StErrorCaptcha    = st(8, "认证码错误")
	StErrorUnknown    = st(9, "未知错误")
	StErrorRank       = st(10, "用户等级错误")

	StErrorSp    = st(11, "逻辑错误")
	StVerifyFail = st(12, "数据认证失败")

	//1~2 AppId 3~4 状态
	StUsernameExist    = st(1000, "用户名已存在")
	StUsernameNotExist = st(1001, "用户名不存在")
	StEmailExist       = st(1002, "邮箱已存在")
	StEmailNotExist    = st(1003, "邮箱不存在")
	StUserPassError    = st(1004, "密码错误")
	StUserPassErrorMax = st(1005, "密码错误")
	StUserDisable      = st(1006, "用户禁用")
	//--- 1000~1099 base-user

)

func st(Code int32, Desc string) *ST {
	return &ST{Code: Code, Msg: Desc}
}

type ST struct {
	Code int32
	Msg  string
}

func (c *ST) ResultNil() *Result {
	return &Result{Status: &ST{Code: c.Code, Msg: c.Msg}}
}
func (c *ST) Result(result string) *Result {
	return &Result{Result: result, Status: &ST{Code: c.Code, Msg: c.Msg}}
}

func (c *ST) OutJSON(gin *gin.Context, result interface{}) {
	gin.Header("Content-Type", "application/json; charset=utf-8")
	gin.JSON(http.StatusOK, map[string]interface{}{"Result": result, "Status": c})
}

func (c *ST) To(m map[string]interface{}) {
	m["Status"] = c
	//return &Result{Status: &ST{Code: c.Code, Msg: c.Msg}}
}
