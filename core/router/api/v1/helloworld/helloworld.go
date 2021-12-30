package helloworld

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/utils/resp"
)

func GetHelloWorld(c *gin.Context) {
	testList := []int{1, 2, 3}
	fmt.Println(testList[4])
	resp.JSON(c, resp.Success, "", "helloWorld")
}
