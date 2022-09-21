package pkg

const (
	ERROR_OK                      = 10000 //完成
	ERROR_JSONPARSE               = 10001 //JSON解析失败
	ERROR_SQL                     = 10002 //SQL执行失败
	ERROR_PARAM                   = 10003 //参数错误
	ERROR_DATA_NOT_FUOUND         = 10004 //数据不存在
	ERROR_DATA_EXIST              = 10005 //数据已存在
	ERROR_TOTKEN_NULL             = 10006 //token为空
	ERROR_TOTKEN                  = 10007 //token校验失败
	ERROR_TOKEN_TIMEOUT           = 10008 //token超时
	ERROR_USERNAME_PASSWORD_ERROR = 10009 //用户名或者密码不一致
	ERROR_TOKEN_GENERATE_ERROR    = 10010 //Token生成失败
)

var (
	ErrorMsg = map[int]string{
		ERROR_OK:                      "Success",
		ERROR_JSONPARSE:               "Json Parse Error",
		ERROR_SQL:                     "SQL Execute Error",
		ERROR_PARAM:                   "Param Error",
		ERROR_DATA_NOT_FUOUND:         "Data Not Found",
		ERROR_DATA_EXIST:              "Data Exist",
		ERROR_TOTKEN:                  "Token Verify Error",
		ERROR_TOTKEN_NULL:             "Token Is Null",
		ERROR_TOKEN_TIMEOUT:           "Token Is Timeout",
		ERROR_USERNAME_PASSWORD_ERROR: "UserName And PassWord Verify Failed",
		ERROR_TOKEN_GENERATE_ERROR:    "Token Generate Failed",
	}
)

func GetErrorMsg(code int) string {
	return ErrorMsg[code]
}
