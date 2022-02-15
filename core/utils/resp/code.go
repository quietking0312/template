package resp

// 错误码,前端对1-599 进行返回拦截并打印错误信息, 600 及以上打印错误信息并返回。 业务错误码建议设定为600及以上
const (
	Success        = 0   // 成功
	ErrArgs        = 300 // 参数错误
	ErrServer      = 500 // 服务错误
	ErrTokenExpire = 501 // token expire
	ErrPermission  = 505
)

var msgCode = map[int]string{
	Success:        "ok",
	ErrTokenExpire: "token is expire",
	ErrPermission:  "权限不足",
}

func GetMsg(code int) string {
	if msg, exists := msgCode[code]; exists {
		return msg
	}
	return ""
}
