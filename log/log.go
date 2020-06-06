package log

import (
	"github.com/crazychat/blog-gin/common"
	"github.com/crazychat/blog-gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var errFile *os.File
var ErrLogger *log.Logger
const filePath = "blogErr.log"

// 日志文件
func InitLog() {
	// 创建日志文件
	var err error
	errFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil{
		log.Fatalln("failed to open file: ", err)
	}
	ErrLogger = log.New(io.MultiWriter(os.Stderr,errFile),"",log.Ldate | log.Ltime | log.Lshortfile)
}

// RecordLog 记录日志
func RecordLog(c *gin.Context, err error) {
	str := c.Request.Method + " | " + c.FullPath() + " | Err: " + err.Error()
	ErrLogger.Println(str)
}

func RecordErr(err error) {
	ErrLogger.Println(time.Now(), "Err: ", err)
}

// GetErrorLog 获取博客错误日志文件内容
func GetErrorLog(c *gin.Context) {
	fi, err := os.Open(filePath)
	if err != nil {
		common.RejectResult(c, models.BACKERROR, "")
		return
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		common.RejectResult(c, models.BACKERROR, "")
		return
	}
	common.ResolveResult(c, models.CONTROLLER_SUCCESS, string(fd))
}

func CloseLogFile() {
	if err := errFile.Close(); err != nil {
		ErrLogger.Fatal("日志文件无法关闭，请检查")
	}
}