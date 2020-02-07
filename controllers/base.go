package controllers

import (
	"blog-gin/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetID struct {
	ID uint32 `json:"id"`
}

// ResolveResult 成功, 返回成功信息
func ResolveResult(c *gin.Context, msg string, data interface{}) {
	c.JSON(httpCode(msg), gin.H{ "data": data })
}

// RejectResult 失败, 返回错误信息
func RejectResult(c *gin.Context, msg string) {
	c.JSON(httpCode(msg), gin.H{ "msg": msg })
}

// 获取http状态码
func httpCode(msg string) int {
	switch msg {
	case models.ANALYSIS_ERROR, models.NO_NULL, models.NAMEERROR, models.PASSWORDERROR, models.FAILEDERROR:
		return 400
	case models.NO_EXIST:
		return 404
	case models.EXISTED, models.CONTROLLER_SUCCESS:
		return 200
	case models.NO_POWER:
		return 401
	case models.BACKERROR:
		return 500
	default:
		return 400
	}
}

// AnalysisJSON 解析JSON数据
func AnalysisJSON(c *gin.Context, obj interface{}) bool {
	err := c.ShouldBindJSON(obj)
	if !CheckErr(err) {
		RejectResult(c, models.ANALYSIS_ERROR)
		return false
	}
	return true
}

// CheckErr 统一处理错误
func CheckErr(err error) bool {
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return true
}