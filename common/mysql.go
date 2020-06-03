package common

const (
	mysqlUser     string = "root"
	mysqlPassword string = "aini1314584"
	mysqlHost     string = "127.0.0.1:3306"
	mysqlDatabase string = "blog"
)

var DB *sqlx.DB

// ConnectMysql 连接mysql数据库
func ConnectMysql() error {
	// 获取配置文件数据
	par := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ")/" + mysqlDatabase
	mysqlDB, _ := sqlx.Open("mysql", par)
	if err := mysqlDB.Ping(); err != nil {
		return err
	}
	return nil
}