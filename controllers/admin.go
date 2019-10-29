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
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, 400, ANALYSIS_ERROR)
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

// PutAdmin 修改管理员信息
func PutAdmin(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, 400, ANALYSIS_ERROR)
	}
	msg, err := models.SetAdmin(admin)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, admin)
}

// PutAdminPwd 修改管理员密码
func PutAdminPwd(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, 400, ANALYSIS_ERROR)
	}
	msg, err := models.SetPassword(admin.Password)
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, "修改成功")
}

// AddView 添加访问量
func AddView(c *gin.Context) {
	msg, err := models.AddView()
	if !CheckErr(err) {
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, msg)
}
