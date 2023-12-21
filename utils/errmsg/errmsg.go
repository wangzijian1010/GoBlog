package errmsg

// 同级目录下不允许有两个不同名称的包

const (
	SUCCSE = 200
	ERROR  = 500

	// code = 1000 表示为用户模块的错误
	ERROR_USERNAME_USED      = 1001 //用户名重名
	ERROR_PASSWORD_WRONG     = 1002 // 密码错误
	ERROR_USERNAME_NOT_EXIST = 1003 // 用户名不存在
	ERROR_TOKEN_EXIST        = 1004 // JWT中TOKEN不存在
	ERROR_TOKEN_RUNTIME      = 1005 // 运行时TOKEN报错
	ERROR_TOKEN_WORNG        = 1006 // TOKEN报错
	ERROR_TOKEN_TYPE_WORNG   = 1007
	ERROR_USER_NO_RIGHT      = 1008
	// code = 2000 开头为文章模块的错误
	ERROR_CATENAME_USED = 2001

	// code = 3000 开头的时候为分类模块的错误
	ERROR_ARTICLE_NOT_FOUND = 3001
	ERROR_CATE_NOT_FOUND    = 3002
)

var codemsg = map[int]string{
	SUCCSE:                   "OK",
	ERROR:                    "Fail",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USERNAME_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:        "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过期",
	ERROR_TOKEN_WORNG:        "TOKEN不存在",
	ERROR_TOKEN_TYPE_WORNG:   "TOKEN格式错误",
	ERROR_CATENAME_USED:      "该分类已存在",
	ERROR_ARTICLE_NOT_FOUND:  "文章不存在",
	ERROR_CATE_NOT_FOUND:     "该分类不存在",
	ERROR_USER_NO_RIGHT:      "该用户无权限",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
