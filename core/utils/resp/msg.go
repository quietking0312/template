package resp

var msgCode = map[int]string{
	Success: "ok",
}

func GetMsg(code int) string {
	if msg, exists := msgCode[code]; exists {
		return msg
	}
	return ""
}
