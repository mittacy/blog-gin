package controller

import (
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// GetErrorLog 获取博客错误日志文件内容
func GetErrorLog(c *gin.Context) {
	fi := log.ErrFile
	if fi == nil {
		if err := log.InitLog(); err != nil {
			common.RejectResult(c, models.BACKERROR, "")
			return
		}
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		common.RejectResult(c, models.BACKERROR, "")
		return
	}
	common.ResolveResult(c, models.CONTROLLER_SUCCESS, string(fd))
}
