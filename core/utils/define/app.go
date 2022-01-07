package define

import "server/common/cryptos"

const TokenKey = "auth"

const (
	UserStateOn     = 1 // 可用
	UserStateOff    = 2 // 不可用
	UserStateDelete = 4 // 删除, 标记为删除后将不会展示给前端
)

func CryptosPass(password string) string {
	return cryptos.Get32MD5(password)
}
