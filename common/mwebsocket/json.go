package mwebsocket

import (
	"encoding/json"
	"fmt"
)

type JSONPack struct {
}

func (p *JSONPack) ReadPack(message []byte) {
	var m map[string]interface{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		return
	}
	fmt.Println(m)
}

func (p *JSONPack) WritePack(args interface{}) []byte {
	data, err := json.Marshal(args)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}
