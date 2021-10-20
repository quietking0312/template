package mwebsocket

import (
	"fmt"
)

type HandlerFunc func()

type WSRouter struct {
	mapFunc map[string]HandlerFunc
}

func NewRouter() *WSRouter {
	return &WSRouter{
		mapFunc: make(map[string]HandlerFunc),
	}
}

func (w *WSRouter) GoFunc(c *WSConn, request interface{}) {
	path := "null"
	fmt.Println("====")
	fmt.Println(request)
	if _, ok := w.mapFunc[path]; ok {
		go w.mapFunc[path]()
	}
}
