package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ANALYSIS_ERROR string = "JSON解析错误"
)

// ResolveResult 成功, 返回成功信息
func ResolveResult(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
	})
}

// RejectResult 失败, 返回错误信息
func RejectResult(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"success": false,
		"msg":     msg,
	})
}

// CheckErr 检查错误
func CheckErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}