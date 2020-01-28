package controllers

import (
	"blog-gin/models"
	"github.com/gin-gonic/gin"
)

// GetAdmin 获取管理员信息
func GetAdmin(c *gin.Context) {
	admin, msg, err := models.GetAdmin()
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, admin)
}

// PostAdmin 登录管理员
func PostAdmin(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, models.ANALYSIS_ERROR)
		return
	}
	// 验证是否正确
	msg, err := models.IsRightAdmin(admin)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	// 登录成功, 生成token
	tokenStr, err := CreateToken(admin.Name)
	if !CheckErr(err) {
		RejectResult(c, models.CREATETOKENERROR)
		return
	}
	ResolveResult(c, models.CONTROLLER_SUCCESS, tokenStr)
}

// PutAdmin 修改管理员信息
func PutAdmin(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, models.ANALYSIS_ERROR)
	}
	msg, err := models.SetAdmin(admin)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, admin)
}

// PutAdminPwd 修改管理员密码
func PutAdminPwd(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, models.ANALYSIS_ERROR)
	}
	msg, err := models.SetPassword(admin.Password)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, nil)
}

// AddView 添加访问量
func AddAdminView(c *gin.Context) {
	msg, err := models.AddAdminView()
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, nil)
}
