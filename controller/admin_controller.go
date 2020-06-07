package controller

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/config"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
	"github.com/crazychat/blog-gin/service"
	"github.com/crazychat/blog-gin/utiles"
	"github.com/gin-gonic/gin"
)

type IAdminController interface {
	Post(c *gin.Context)
	Put(c *gin.Context)
	PutPassword(c *gin.Context)
	Get(c *gin.Context)
	Verify(c *gin.Context)
	InitAdmin() error
}

type AdminController struct {
	AdminService service.IAdminService
}
// NewAdminController 获取Admin控制器
func NewAdminController() IAdminController {
	repo := repository.NewAdminRepository("admin")
	adminService := service.NewAdminService(repo)
	return &AdminController{adminService}
}
// InitAdmin 初始化Admin
func (ac *AdminController) InitAdmin() error {
	return ac.AdminService.CreateAdmin(&config.InitAdmin)
}
// Post 登录管理员
func (ac *AdminController) Post(c *gin.Context) {
	ip := c.ClientIP()
	// 1. 检查ip访问次数是否超过
	if times, err := cache.GetIPTimes(ip); err != nil || times > config.IPPostTimes {
		common.RejectResult(c, common.LOGINFREQUENTLY, &model.Admin{})
		return
	}
	// 不超过，可以访问，增加ip访问记录
	if err := cache.IPIncr(ip); err != nil {
		log.RecordLog(c, err)
	}
	// 2. 解析json数据到结构体admin
	admin := &model.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Admin{})
		return
	}
	// 3. 验证用户名密码是否正确
	adminRight, err := ac.AdminService.GetAdmin()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 用户名或密码错误
	if admin.Name != adminRight.Name {
		common.RejectResult(c, common.NAMEERROR, &model.Admin{})
		return
	} else if utiles.Encryption(admin.Password) != adminRight.Password {
		common.RejectResult(c, common.PASSWORDERROR, &model.Admin{})
		return
	}
	// 4. 正确 -> 删除ip错误记录，生成 token 返回
	if err := cache.IPDel(ip); err != nil {
		log.RecordLog(c, err)
	}
	tokenStr, err := utiles.CreateToken(admin.Password)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 5. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, tokenStr)
	// 6. 缓存token
	cache.SetToken(tokenStr)
}
// Put 修改管理员信息
func (ac *AdminController) Put(c *gin.Context) {
	// 1. 解析json数据到结构体
	admin := &model.Admin{}
	if err := c.ShouldBindJSON(admin); err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Admin{})
		return
	}
	// 2. 操作数据库
	err := ac.AdminService.UpdateAdminInfo(admin)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, admin)
}
// PutPassword 修改管理员密码
func (ac *AdminController) PutPassword(c *gin.Context) {
	// 1. 解析json数据到结构体admin
	admin := &model.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Admin{})
		log.RecordLog(c, err)
		return
	}
	// 2. 修改
	admin.Password = utiles.Encryption(admin.Password)
	err := ac.AdminService.UpdateAdminPassword(admin)
	if err != nil {
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		log.RecordLog(c, err)
		return
	}
	// 3. 返回结果
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, &model.Admin{})
}
// Get 获取管理员信息
func (ac *AdminController) Get(c *gin.Context) {
	// 1. 从缓存器或数据库取数据（不包含密码）
	admin, err := ac.AdminService.GetAdminInfo()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 2. 返回前端
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, admin)
}
// Verify 验证登录
func (ac *AdminController) Verify(c *gin.Context) {
	// 1. 获取请求token
	adminToken := c.Request.Header.Get(cache.TokenName)
	if adminToken == "" {
		common.RejectResult(c, common.NO_POWER, &model.Admin{})
		return
	}
	// 2. 获取缓存器token
	token, isExist := cache.GetToken()
	if !isExist || adminToken != token {
		common.RejectResult(c, common.NO_POWER, &model.Admin{})
		return
	}
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, nil)
}

