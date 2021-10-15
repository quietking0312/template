package helloworld

import (
	"github.com/gin-gonic/gin"
	"server/core/utils/resp"
)

func GetHelloWorld(c *gin.Context) {

	resp.JSON(c, resp.Success, "", "helloWorld")
}
