package controllers

import (
	"blog-gin/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

// GetArticleID 获取首页文章id
func GetArticleID(c *gin.Context) {
	admin, msg, err := models.GetArticleID()
	if !CheckErr(err) {
		fmt.Println(err)
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, admin)
}

// PostAdmin 登录管理员
func PostAdmin(c *gin.Context) {
	fmt.Println("IP: -> ", c.ClientIP())
	if !CheckIP(c) {
		RejectResult(c, models.NO_POWER)
		return
	}
	admin := &models.Admin{}
	// 解析json数据到结构体admin
	if err := c.ShouldBindJSON(admin); !CheckErr(err) {
		RejectResult(c, models.ANALYSIS_ERROR)
		return
	}
	// 验证是否正确
	msg, err := models.IsRightAdmin(admin)
	if !CheckErr(err) {
		if msg == models.NAMEERROR || msg == models.PASSWORDERROR {
			AddErrorIP(c)
		}
		RejectResult(c, msg)
		return
	}
	// 登录成功, 生成token
	DelIP(c)
	tokenStr, err := CreateToken(admin.Name)
	if !CheckErr(err) {
		RejectResult(c, models.FAILEDERROR)
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

// Verify 验证登录
func Verify(c *gin.Context) {
	tokenStr := c.Request.Header.Get(tokenName)
	// token是否存在
	if tokenStr == "" {
		RejectResult(c, models.NO_POWER)
		return
	}
	// 解析token
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			RejectResult(c, models.NO_POWER)
			return nil, fmt.Errorf(models.NO_POWER)
		}
		return []byte(serectKey), nil
	})
	// token是否过期
	if !token.Valid {
		RejectResult(c, models.NO_POWER)
		return
	}
	ResolveResult(c, models.CONTROLLER_SUCCESS, nil)
}

// PutArticleID 修改主页文章id
func PutArticleID(c *gin.Context) {
	getID := GetID{}
	if !AnalysisJSON(c, &getID) {
		return
	}
	msg, err := models.SetArticleID(getID.ID)
	if !CheckErr(err) {
		RejectResult(c, msg)
		return
	}
	ResolveResult(c, msg, msg)
}