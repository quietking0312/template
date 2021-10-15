package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/core/utils/define"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", strings.Join([]string{"Content-Type", define.TokenKey}, ","))
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, ws")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
