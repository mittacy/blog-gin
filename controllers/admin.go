package controllers

import (
	"github.com/crazychat/blog-gin/models"
	"github.com/gin-gonic/gin"
)

// GetAdmin 获取管理员信息
func GetAdmin(c *gin.Context) {
	admin, msg, err := models.GetAdmin()
	if !CheckErr(err, c) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, admin)
}

// PostAdmin 登录管理员
func PostAdmin(c *gin.Context) {
	ip := c.ClientIP()
	// 检查ip访问次数是否超过
	if !models.CheckIPRequestTimes(ip) {
		RejectResult(c, models.LOGINFREQUENTLY)
		return
	}
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if !AnalysisJSON(c ,admin) {
		return
	}
	// 验证是否正确
	msg, err := models.CheckPassword(admin)
	if err != nil {
		if msg == models.NAMEERROR || msg == models.PASSWORDERROR {
			// 错误ip请求次数+1
			CheckErr(models.IncrIP(ip), c)
		}
		RejectResult(c, msg)
		return
	}
	// 登录成功, 生成token
	CheckErr(models.DelIP(ip), c)
	tokenStr, err := CreateToken()
	if !CheckErr(err, c) {
		RejectResult(c, models.FAILEDERROR)
		return
	}
	CheckErr(models.SaveToken(tokenStr), c)
	ResolveResult(c, models.CONTROLLER_SUCCESS, tokenStr)
}

// PutAdmin 修改管理员信息
func PutAdmin(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err, c) {
		RejectResult(c, models.ANALYSIS_ERROR)
	}
	msg, err := models.SetAdmin(admin)
	if !CheckErr(err, c) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, admin)
}

// PutAdminPwd 修改管理员密码
func PutAdminPwd(c *gin.Context) {
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err, c) {
		RejectResult(c, models.ANALYSIS_ERROR)
	}
	msg, err := models.SetPassword(admin.Password)
	if !CheckErr(err, c) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, nil)
}

// Verify 验证登录
func Verify(c *gin.Context) {
	adminToken := c.Request.Header.Get(models.TokenName)
	// token是否存在
	if adminToken == "" {
		RejectResult(c, models.NO_POWER)
		return
	}
	// redis比对
	if !models.Verify(adminToken) {
		RejectResult(c, models.NO_POWER)
		return
	}
	ResolveResult(c, models.CONTROLLER_SUCCESS, nil)
}