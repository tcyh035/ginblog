package errmsg

const (
	// Success 成功error code
	Success = 200
	// Error 失败error code
	Error = 500

	// ErrorUserNameUsed 用户名重复
	ErrorUserNameUsed = 1001
	// ErrorPasswordWrong 密码错误
	ErrorPasswordWrong = 1002
	// ErrorUserNotExist 用户不存在
	ErrorUserNotExist = 1003
	// ErrorTokenNotExist Token不存在
	ErrorTokenNotExist = 1004
	// ErrorTokenRuntime Token超时
	ErrorTokenRuntime = 1005
	// ErrorTokenWrong Token验证失败
	ErrorTokenWrong = 1006
	// ErrorTokenTypeWrong Token格式错误
	ErrorTokenTypeWrong = 10007
	// ErrorUserNoRight 用户没权限
	ErrorUserNoRight = 10008

	// ErrorCategoryUsed 分类已被使用
	ErrorCategoryUsed = 2001
	// ErrorCategoryNotExist 分类不存在
	ErrorCategoryNotExist = 2002

	// ErrorArticleNotExist 文章不存在
	ErrorArticleNotExist = 3001
)

var codemsg = map[int]string{
	Success:               "OK",
	Error:                 "FAIL",
	ErrorUserNameUsed:     "用户名已存在",
	ErrorPasswordWrong:    "密码错误",
	ErrorUserNotExist:     "用户不存在",
	ErrorTokenNotExist:    "Token不存在",
	ErrorTokenRuntime:     "Token已过期",
	ErrorTokenWrong:       "Token不正确",
	ErrorTokenTypeWrong:   "Token格式错误",
	ErrorCategoryUsed:     "分类已存在",
	ErrorCategoryNotExist: "分类不存在",
	ErrorArticleNotExist:  "文章不存在",
	ErrorUserNoRight:      "用户没权限",
}

// GetErrorMessage 得到错误信息
func GetErrorMessage(code int) string {
	return codemsg[code]
}
