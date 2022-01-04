package resp

const (
	Success   = 0   // 成功
	ErrArgs   = 300 // 参数错误
	ErrServer = 500 // 服务错误
)

var msgCode = map[int]string{
	Success: "ok",
}

func GetMsg(code int) string {
	if msg, exists := msgCode[code]; exists {
		return msg
	}
	return ""
}
