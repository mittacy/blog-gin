package controller

import (
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/models"
	"github.com/crazychat/blog-gin/repository"
	"github.com/crazychat/blog-gin/service"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	AdminService service.IAdminService
}
// GetAdminController 获取Admin控制器
func GetAdminController() *AdminController {
	repo := repository.NewAdminRepository("admin")
	adminService := service.NewAdminService(repo)
	return &AdminController{adminService}
}

// GetAdmin 获取管理员信息
func (ac *AdminController) GetAdmin(c *gin.Context) {
	admin, err := ac.AdminService.GetAdminInfo()
	if err != nil {
		common.RejectResult(c, common.BACKERROR, admin)
	}
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, admin)
}

// PostAdmin 登录管理员
func (ac *AdminController) PostAdmin(c *gin.Context) {
	ip := c.ClientIP()
	// 1. 检查ip访问次数是否超过
	if !common.CheckIPRequestTimes(ip) {
		common.RejectResult(c, common.LOGINFREQUENTLY, &model.Admin{})
		return
	}
	admin := &models.Admin{}
	// 2. 解析json数据到结构体admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		return
	}
	// 3. 验证用户名密码是否正确
	pwd, err := ac.AdminService.GetAdminPassword()
	if err != nil {
		common.RejectResult(c, common.BACKERROR, &model.Admin{})
	}
	// 用户名密码错误 -> 记录次数并返回
	admin.Password = common.Encryption(admin.Password)
	if admin.Password != pwd {
		if err := common.IncrIP(ip); err != nil {
			// todo 处理err
		}
		common.RejectResult(c, common.PASSWORDERROR, &model.Admin{})
		return
	}
	// 4. 正确 -> 删除ip错误记录，生成 token 返回
	if err := common.DelIP(ip); err != nil {
		// todo 处理err
	}
	//tokenStr, err := common.CreateToken(admin.Password)
	//if err != nil {
	//	// todo 处理err
	//	common.RejectResult(c, common.BACKERROR, &model.Admin{})
	//	return
	//}
	// todo 缓存token
}
