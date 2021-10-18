package helloworld

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/utils/resp"
)

func GetHelloWorld(c *gin.Context) {
	a := []int{1}
	fmt.Println(a[2])
	resp.JSON(c, resp.Success, "", "helloWorld")
}
