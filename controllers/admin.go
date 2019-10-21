package controllers

import (
	"blog-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAdmin 获取管理员信息
func GetAdmin(c *gin.Context) {
	admin, msg, err := models.GetAdmin()
	if CheckErr(err) {
		ResolveResult(c, http.StatusOK, admin)
		return
	}
	RejectResult(c, http.StatusBadRequest, msg)
}

// PostAdmin 登录管理员
func PostAdmin(c *gin.Context) {
	admin := &models.Admin{}
	err := c.ShouldBindJSON(admin)
	if !CheckErr(err) {
		RejectResult(c, http.StatusBadRequest, ANALYSIS_ERROR)
		return
	}
	// 验证是否正确
	msg, err := models.IsRightAdmin(admin)
	if CheckErr(err) {
		ResolveResult(c, http.StatusOK, msg)
		return
	}
	RejectResult(c, http.StatusBadRequest, msg)
}
