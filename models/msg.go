package models

// 常见错误定义
const (
	NO_EXIST          string = "资源不存在"
	EXISTED           string = "数据已存在"
	CATE_NO_EXIST     string = "分类不存在"
	CATE_EXIST        string = "分类存在"
	ARTICLE_EXIST     string = "文章存在"
	ARTICLE_NO_EXIST  string = "文章不存在"
	ADMIN_NO_EXIST    string = "管理员不存在"
	NO_POWER          string = "权限不足"
	SQL_ERROR         string = "数据库操作错误"
	CONVERSIOIN_ERROR string = "数据转换出错"
	ANALYSIS_ERROR    string = "JSON解析错误"
	COMMITNONE        string = "提交数据为空"
)
