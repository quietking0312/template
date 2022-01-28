package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(c *gin.Context, code int, message string, data interface{}) {
	var resp Response

	resp = Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	if resp.Message == "" {
		resp.Message = GetMsg(code)
	}

	c.JSON(http.StatusOK, resp)
	defer abort(c)
	return
}

func String(c *gin.Context, message string) {
	c.Writer.WriteString(message)
	defer abort(c)
	return
}

func File(c *gin.Context, filePath string) {
	c.File(filePath)
	defer abort(c)
	return
}

func abort(c *gin.Context) {
	c.Abort()
}
