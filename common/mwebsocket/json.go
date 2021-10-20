package mwebsocket

import (
	"encoding/json"
	"fmt"
)

type JSONPack struct{}

func (p *JSONPack) ReadPack(message []byte) interface{} {
	var m map[string]interface{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return m
}

func (p *JSONPack) WritePack(args interface{}) ([]byte, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	return data, nil
}
