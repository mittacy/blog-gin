package controllers

import (
	"blog-gin/models"

	"github.com/gin-gonic/gin"
)

// GetAdmin 获取管理员信息
func GetAdmin(c *gin.Context) {
	admin, msg, err := models.GetAdmin()
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, admin)
}

// PostAdmin 登录管理员
func PostAdmin(c *gin.Context) {
	admin := &models.Admin{}
	err := c.ShouldBindJSON(admin)
	if !CheckErr(err) {
		RejectResult(c, 400, ANALYSIS_ERROR)
		return
	}
	// 验证是否正确
	msg, err := models.IsRightAdmin(admin)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	// 登录成功, 生成token
	tokenStr, err := CreateToken(admin.Name)
	if !CheckErr(err) {
		RejectResult(c, 500, "生成token失败")
		return
	}
	ResolveResult(c, 200, tokenStr)
}
