package constant

const (
	AccountEmptyAuthHeader   = "请求头未携带token信息"
	AccountInvalidAuthHeader = "请求头格式不合法"
	AccountInvalidToken      = "Token不合法"
	AccountTokenExpired      = "Token已过期"

	AccountAlreadyExist       = "存在相同账号的用户"
	AccountEmailAlreadyExist  = "存在相同邮箱的用户"
	AccountNotExist           = "用户不存在"
	AccountNotCorrectPassword = "用户名/密码组合不正确"
)
