package common

import (
	"github.com/gin-gonic/gin"
)

// ResolveResult 成功, 返回成功信息
func ResolveResult(c *gin.Context, msg string, data interface{}) {
	c.JSON(httpCode(msg), gin.H {
		"code": 1,
		"data": data,
		"msg": msg,
	})
}
// RejectResult 失败, 返回错误信息
func RejectResult(c *gin.Context, msg string, data interface{}) {
	c.JSON(httpCode(msg),  gin.H {
		"code": -1,
		"msg": msg,
		"data": data,
	})
}
// 获取http状态码
func httpCode(msg string) int {
	switch msg {
	//case EXISTED, ANALYSIS_ERROR, NO_NULL, NAMEERROR, PASSWORDERROR, FAILEDERROR:
	//	return 400
	case NO_EXIST:
		return 404
	case CONTROLLER_SUCCESS:
		return 200
	case NO_POWER:
		return 401
	case BACKERROR:
		return 500
	default:
		return 400
	}
}
