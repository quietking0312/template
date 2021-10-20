package mwebsocket

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type WSConn struct {
	sync.RWMutex
	uuid   string
	pack   Pack
	Socket *websocket.Conn
	router Router
}

// Pack 封包解包类
type Pack interface {
	ReadPack(message []byte) interface{}
	WritePack(data interface{}) ([]byte, error)
}

// Router 路由类
type Router interface {
	GoFunc(c *WSConn, request interface{})
}

func NewWSConn(conn *websocket.Conn, pack Pack, router Router) *WSConn {
	c := &WSConn{
		uuid:   uuid.New().String(),
		pack:   pack,
		Socket: conn,
		router: router,
	}
	go c.read()
	return c
}

func (w *WSConn) read() {
	defer func() {
		_ = w.Close()
	}()
	for {
		if w.Socket != nil {
			_, message, err := w.Socket.ReadMessage()
			if err != nil {

				return
			}
			// 解包
			data := w.pack.ReadPack(message)
			if data != nil {
				// 传递链接及解包信息给路由接口
				w.router.GoFunc(w, data)
			}
		} else {
			break
		}
	}
}

func (w *WSConn) Write(res interface{}) error {
	w.Lock()
	defer w.Unlock()
	// 封包
	data, err := w.pack.WritePack(res)
	if err != nil {
		return err
	}
	if w.Socket != nil {
		return w.Socket.WriteMessage(websocket.BinaryMessage, data)
	} else {
		return errors.New("conn is null")
	}
}

func (w *WSConn) Close() error {
	w.Lock()
	defer func() {
		w.Socket = nil
		w.Unlock()
	}()
	return w.Socket.Close()
}
