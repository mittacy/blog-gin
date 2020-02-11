package controllers

import (
	"blog-gin/models"
	"github.com/gin-gonic/gin"
)

type GetID struct {
	ID uint32 `json:"id"`
}

var Ip = map[string]int{}


// ResolveResult 成功, 返回成功信息
func ResolveResult(c *gin.Context, msg string, data interface{}) {
	c.JSON(httpCode(msg), gin.H{ "data": data, "msg": msg })
}

// RejectResult 失败, 返回错误信息
func RejectResult(c *gin.Context, msg string) {
	c.JSON(httpCode(msg), gin.H{ "msg": msg })
}

// 获取http状态码
func httpCode(msg string) int {
	switch msg {
	case models.EXISTED, models.ANALYSIS_ERROR, models.NO_NULL, models.NAMEERROR, models.PASSWORDERROR, models.FAILEDERROR:
		return 400
	case models.NO_EXIST:
		return 404
	case models.CONTROLLER_SUCCESS:
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
		return false
	}
	return true
}

// CheckIP 检查ip错误次数
func CheckIP(c *gin.Context) bool {
	ip := c.ClientIP()
	if num, exists := Ip[ip]; exists && num >= 5 {
		return false
	}
	return true
}
// AddErrorIP 增加错误记录ip
func AddErrorIP(c *gin.Context) {
	// 保存ip错误次数，超过五次封锁该ip
	ip := c.ClientIP()
	num, exists := Ip[ip]
	if exists {
		Ip[ip] = num + 1
		return
	}
	Ip[ip] = 1
	return
}
// DelIP 删除ip记录
func DelIP(c *gin.Context) {
	ip := c.ClientIP()
	if _, exists := Ip[ip]; exists {
		delete(Ip, ip)
	}
	return
}