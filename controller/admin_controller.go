package controller

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/models"
	"github.com/crazychat/blog-gin/repository"
	"github.com/crazychat/blog-gin/service"
	"github.com/gin-gonic/gin"
)

type IAdminController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	PutPassword(c *gin.Context)
	Verify(c *gin.Context)
}

type AdminController struct {
	AdminService service.IAdminService
}
// GetAdminController 获取Admin控制器
func GetAdminController() IAdminController {
	repo := repository.NewAdminRepository("admin")
	adminService := service.NewAdminService(repo)
	return &AdminController{adminService}
}

// Get 获取管理员信息
func (ac *AdminController) Get(c *gin.Context) {
	// 1. 取数据
	admin, err := ac.AdminService.GetAdminInfo()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, admin)
	}
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, admin)
}
// Post 登录管理员
func (ac *AdminController) Post(c *gin.Context) {
	ip := c.ClientIP()
	// 1. 检查ip访问次数是否超过
	if !common.CheckIPRequestPower(ip) {
		common.RejectResult(c, common.LOGINFREQUENTLY, &model.Admin{})
		return
	}
	// 不超过，可以访问，增加ip访问记录
	if err := common.IncrIP(ip); err != nil {
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
	pwd, err := ac.AdminService.GetAdminPassword()
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 用户名密码错误
	admin.Password = common.Encryption(admin.Password)
	if admin.Password != pwd {
		common.RejectResult(c, common.PASSWORDERROR, &model.Admin{})
		return
	}
	// 4. 正确 -> 删除ip错误记录，生成 token 返回
	if err := common.DelIP(ip); err != nil {
		log.RecordLog(c, err)
	}
	tokenStr, err := common.CreateToken(admin.Password)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		return
	}
	// 5. 缓存token
	common.ResolveResult(c, models.CONTROLLER_SUCCESS, tokenStr)
	common.SaveToken(tokenStr)
}
// Put 修改管理员信息
func (ac *AdminController) Put(c *gin.Context) {
	// 1. 解析json数据到结构体admin
	admin := &model.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.RecordLog(c, err)
		return
	}
	// 2. 修改
	err := ac.AdminService.UpdateAdminInfo(admin)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.ANALYSIS_ERROR, &model.Admin{})
		return
	}
	// 3. 更新缓存器
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, admin)
	cache.UpdateAdminInfo(admin)
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
	err := ac.AdminService.UpdatePassword(admin)
	if err != nil {
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
		log.RecordLog(c, err)
		return
	}
	// 3. 更新缓存器
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, &model.Admin{})
	cache.UpdateAdminInfo(admin)
}
// Verify 验证登录
func (ac *AdminController) Verify(c *gin.Context) {
	// 1. 获取请求token
	adminToken := c.Request.Header.Get(models.TokenName)
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
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, nil)
}

