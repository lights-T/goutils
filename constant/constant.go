package constant

//时间相关
const (
	DateLayoutYear    = "2006"
	DateLayoutMouth   = "2006-01"
	DateLayout        = "2006-01-02"
	DatetimeLayout    = "2006-01-02 15:04:05"
	DatetimeLayoutNa  = "2006-01-02 15:04:05.999"
	DefaultMinTime    = "1971-01-01 00:00:00"
	DefaultMaxTime    = "2099-01-01 00:00:00"
	DatetimeLayoutMin = "2006-01-02 15:04"

	DateDayTime  = int64(86400) //一天的秒
	DateHourTime = int64(3600)  //一小时的秒
	DateMinTime  = int64(60)    //每分钟的秒
)

//错误相关
const (
	ErrUploadParamsIsNotExist = "上传信息不存在"
	ErrUploadFileIsNotExist   = "上传文件不存在"
)
