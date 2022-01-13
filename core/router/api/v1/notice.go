package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"path"
	"server/core/utils/resp"
)

func Notice(c *gin.Context) {
	file, err := ioutil.ReadFile(path.Join("notice.md"))
	if err != nil {
		fmt.Println(err)
	}
	resp.JSON(c, resp.Success, "", string(blackfriday.Run(file, blackfriday.WithExtensions(blackfriday.CommonExtensions))))
}
