package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/utils/resp"
	"time"
)

func GetHelloWorld(c *gin.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, "秒")
	}
	resp.JSON(c, resp.Success, "", "helloWorld")
}
