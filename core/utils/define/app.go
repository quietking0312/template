package define

import (
	"github.com/gin-gonic/gin"
	"server/common/cryptos"
)

const TokenKey = "auth"

const (
	UserStateOn     = 1 // 可用
	UserStateOff    = 2 // 不可用
	UserStateDelete = 4 // 删除, 标记为删除后将不会展示给前端
)

const (
	DefaultPass = "123456"
	AdminPid    = 100000
)

func CryptosPass(password string) string {
	return cryptos.Get32MD5(password)
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get(TokenKey)
	if token == "" {
		token, _ = c.Cookie(TokenKey)
	}
	return token
}
