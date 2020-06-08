package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/blog-gin/common"
	"github.com/mittacy/blog-gin/log"
	"io/ioutil"
	"os"
)

// GetErrorLog 获取博客错误日志文件内容
func GetErrorLog(c *gin.Context) {
	fi, err := os.Open(log.FilePath)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, "")
		return
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		log.RecordLog(c, err)
		common.RejectResult(c, common.BACKERROR, "")
		return
	}
	common.ResolveResult(c, common.CONTROLLER_SUCCESS, string(fd))
}
