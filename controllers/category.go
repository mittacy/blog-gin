package controllers

import (
	"blog-gin/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	cate := models.Category{}
	if !AnalysisJSON(c, &cate) {
		return
	}
	if msg, err := models.CreateCate(&cate); err != nil {
		fmt.Println("插入cate err:", err)
		RejectResult(c, 400, msg)
		return
	}
	ResolveResult(c, 200, cate)
}
