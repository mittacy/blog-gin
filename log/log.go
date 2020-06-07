package log

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

var errFile *os.File
var ErrLogger *log.Logger
const FilePath = "./log/blogErr.log"

// 日志文件
func InitLog() error {
	// 创建日志文件
	var err error
	errFile, err = os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		return err
	}
	ErrLogger = log.New(io.MultiWriter(os.Stderr, errFile),"",log.Ldate | log.Ltime | log.Lshortfile)
	return nil
}

// RecordLog 记录日志
func RecordLog(c *gin.Context, err error) {
	str := c.Request.Method + " | " + c.FullPath() + " | Err: " + err.Error()
	ErrLogger.Println(str)
}

func RecordErr(err error) {
	ErrLogger.Println(err)
}

func CloseLogFile() {
	if err := errFile.Close(); err != nil {
		ErrLogger.Fatal("日志文件无法关闭，请检查")
	}
}