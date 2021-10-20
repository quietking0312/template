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
}

type Pack interface {
	ReadPack(message []byte)
	WritePack(data interface{}) []byte
}

func NewWSConn(conn *websocket.Conn, pack Pack) *WSConn {
	c := &WSConn{
		uuid:   uuid.New().String(),
		pack:   pack,
		Socket: conn,
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
			w.pack.ReadPack(message)
		} else {
			break
		}
	}
}

func (w *WSConn) Write(res interface{}) error {
	w.Lock()
	defer w.Unlock()
	data := w.pack.WritePack(res)
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
