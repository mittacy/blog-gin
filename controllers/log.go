package controllers

import (
	"io"
	"log"
	"os"
)

var errFile *os.File
var ErrLogger *log.Logger

// 日志文件
func init() {
	// 创建日志文件
	var err error
	errFile, err = os.OpenFile("blogErr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil{
		log.Fatalln("failed to open file: ", err)
	}
	ErrLogger = log.New(io.MultiWriter(os.Stderr,errFile),"",log.Ldate | log.Ltime | log.Lshortfile)
}

func CloseLogFile() {
	if err := errFile.Close(); err != nil {
		ErrLogger.Fatal("日志文件无法关闭，请检查")
	}
}
