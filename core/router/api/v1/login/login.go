package login

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"server/core/protocol"
	"server/core/utils/resp"
)

func Login(c *gin.Context) {
	// var req protocol.Login
	// if err := c.ShouldBindWith(&req, binding.ProtoBuf); err != nil {
	// 	fmt.Println(err)
	// 	resp.JSON(c, resp.Success, "", "")
	// 	return
	// }
	// fmt.Println(req)
	var res protocol.Login
	res.Username = "hello"
	res.Password = "world"
	data, err := proto.Marshal(&res)
	if err != nil {
		fmt.Println(err)
	}
	byteBuffer := bytes.NewBuffer(data)
	var x int32
	_ = binary.Read(byteBuffer, binary.BigEndian, &x)

	resp.JSON(c, resp.Success, "", int(x))
}
