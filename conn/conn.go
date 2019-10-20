package conn

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	par := "root:aini1314584@tcp(127.0.0.1:3306)/blog?charset=utf8"
	var err error
	db, err = gorm.Open("mysql", par)
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
		panic("failed to connect database")
	}
	fmt.Println("连接数据库成功")
}

func CreateTables() {

}
