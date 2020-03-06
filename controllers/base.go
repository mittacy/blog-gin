package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/crazychat/blog-gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"time"
)

type GetID struct {
	ID uint32 `json:"id"`
}

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
	if !CheckErr(err, c) {
		RejectResult(c, models.ANALYSIS_ERROR)
		return false
	}
	return true
}

// CheckErr 统一处理错误
func CheckErr(err error, c *gin.Context) bool {
	if err != nil {
		str := c.Request.Method + " | " + c.FullPath() + " | Err: " + err.Error()
		ErrLogger.Println(str)
		return false
	}
	return true
}

// CreateToken 生成token
func CreateToken() (string, error) {
	now := time.Now().Unix()
	h := md5.New()
	_, err := io.WriteString(h, strconv.FormatInt(now, 10))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// CheckAdmin 中间件, 检查权限
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get(models.TokenName)
		// token是否存在
		if tokenStr == "" {
			RejectResult(c, models.NO_POWER)
			c.Abort()
			return
		}
		// 验证token
		if !models.Verify(tokenStr) {
			RejectResult(c, models.NO_POWER)
			c.Abort()
			return
		}
		c.Next()
	}
}