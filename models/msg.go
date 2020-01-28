package models

// 常见错误定义
const (
	UNKONWNERROR       string = "Unkonwn error"
	NO_EXIST           string = "Resource does not exist"
	EXISTED            string = "Resource already exists"
	CATE_NO_EXIST      string = "Category does not exist"
	CATE_EXIST         string = "Category already exists"
	ARTICLE_EXIST      string = "Article already exists"
	ARTICLE_NO_EXIST   string = "Article does not exists"
	ADMIN_NO_EXIST     string = "Administrators does not exist"
	NO_POWER           string = "Permission denied"
	SQL_ERROR          string = "Database operation error"
	CONVERSIOIN_ERROR  string = "Data conversion error"
	ANALYSIS_ERROR     string = "JSON error"
	CHECKCONTENT       string = "请检查提交数据是否正确或分类是否存在"
	NO_NULL            string = "Data cannot be empty"
	CONTROLLER_SUCCESS string = "Success"
	NAMEERROR		   string = "Name error"
	PASSWORDERROR	   string = "Password error"
	CREATETOKENERROR   string = "Create token error"
)
