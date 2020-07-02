package constant

const (
	LoadConfigError     = "读取配置错误"
	LogConfigError      = "配置日志错误"
	DatabaseConfigError = "配置数据库错误"
	AppConfigError      = "配置应用错误"
)

const (
	DatabaseConnectError = "数据库连接出错"
	GetConnectionError   = "获取当前数据库连接出错"
	ORMConfigError       = "配置ORM错误"
	DatabaseMigrateError = "数据库迁移出错"
)

const (
	TransGetTranslatorError = "根据语言选择翻译器出错"
	TransRegisterError      = "国际化出错"
)
