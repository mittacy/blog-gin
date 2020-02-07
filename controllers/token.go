package controllers

import (
	"blog-gin/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	serectKey string = "this is my house"
	tokenName string = "adminToken"
)

// CheckAdmin 中间件, 检查权限
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get(tokenName)
		// token是否存在
		if tokenStr == "" {
			RejectResult(c, models.NO_POWER)
			c.Abort()
			return
		}
		// 解析token
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				RejectResult(c, models.NO_POWER)
				c.Abort()
				return nil, fmt.Errorf(models.NO_POWER)
			}
			return []byte(serectKey), nil
		})
		// token是否过期
		if !token.Valid {
			RejectResult(c, models.NO_POWER)
			c.Abort()
			return
		}
		c.Next()
	}
}

// CreateToken 生成token
func CreateToken(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(), // 可以添加过期时间
		"iat":      time.Now().Unix(),
	})
	return token.SignedString([]byte(serectKey)) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

