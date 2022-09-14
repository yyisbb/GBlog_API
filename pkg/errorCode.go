package pkg

const (
	ERROR_OK              = 10000 //完成
	ERROR_JSONPARSE       = 10001 //JSON解析失败
	ERROR_SQL             = 10002 //SQL执行失败
	ERROR_PARAM           = 10003 //参数错误
	ERROR_DATA_NOT_FUOUND = 10004 //数据不存在
	ERROR_DATA_EXIST      = 10005 //数据已存在
)

var (
	ErrorMsg = map[int]string{
		ERROR_OK:              "Success",
		ERROR_JSONPARSE:       "Json Parse Error",
		ERROR_SQL:             "SQL Execute Error",
		ERROR_PARAM:           "Param Error",
		ERROR_DATA_NOT_FUOUND: "Data Not Found",
		ERROR_DATA_EXIST:      "Data Exist",
	}
)

func GetErrorMsg(code int) string {
	return ErrorMsg[code]
}
