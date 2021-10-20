package mwebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestNewWSConn(t *testing.T) {
	var pack = &JSONPack{}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  65535,
			WriteBufferSize: 65535,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		ws, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		NewWSConn(ws, pack)
	})
	http.ListenAndServe("127.0.0.1:8787", nil)
}

func TestNewWSConn2(t *testing.T) {
	var pack = &JSONPack{}
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8787", Path: "/ws"}
	ws, resp, err := websocket.DefaultDialer.Dial(u.String(), http.Header{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Body)
	conn := NewWSConn(ws, pack)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			for i := 0; i < 10; i++ {
				data := map[string]interface{}{
					"hello": "hello",
					"world": 2,
				}
				go func(i int) {
					data["index"] = i
					if err := conn.Write(data); err != nil {
						return
					}

				}(i)
			}
		}
	}
}
