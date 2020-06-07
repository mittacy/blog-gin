package router

import (
	"fmt"
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
// StaticMiddleware 过滤前端请求
func StaticMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if strings.Index(url, "api") >= 0 {
			c.Next()
		} else {
			// todo 增加博客访问量
			cache.AdminViewCacheIncr()
			c.HTML(200, "index.html", gin.H{"msg": "Success"})
		}
		return
	}
}
// CorsMiddleware 允许跨域
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma, adminToken")
			// 允许跨域设置,可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
// VerifyMiddleware 中间件, 检查权限
func VerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取请求token
		adminToken := c.Request.Header.Get(cache.TokenName)
		if adminToken == "" {
			common.RejectResult(c, common.NO_POWER, &model.Admin{})
			return
		}
		// 2. 获取数据库tokne
		token, isExist := cache.GetToken()
		if !isExist || adminToken != token {
			common.RejectResult(c, common.NO_POWER, &model.Admin{})
			return
		}
		// 3. 验证通过
		c.Next()
	}
}
