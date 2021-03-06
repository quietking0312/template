package mwebsocket

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"math"
	"net/http"
	"net/url"
	"testing"
	"time"
	"unsafe"
)

func TestNewWSConn(t *testing.T) {
	var pack = &JSONPack{}
	var router = NewRouter()
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
		NewWSConn(ws, pack, router)
	})
	http.ListenAndServe("127.0.0.1:8787", nil)
}

func TestNewWSConn2(t *testing.T) {
	var pack = &JSONPack{}
	var router = NewRouter()
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8787", Path: "/ws"}
	ws, resp, err := websocket.DefaultDialer.Dial(u.String(), http.Header{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp.Body)
	conn := NewWSConn(ws, pack, router)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			for i := 0; i < 10; i++ {
				go func(i int) {
					data := map[string]interface{}{
						"hello": "hello",
						"world": 2,
					}
					data["index"] = i
					if err := conn.Write(data); err != nil {
						return
					}

				}(i)
			}
		}
	}
}

func TestNewWSConn3(t *testing.T) {
	var pack = &JSONPack{}
	fmt.Printf("%d\n", unsafe.Sizeof(pack))
	var nul = struct{}{}
	fmt.Printf("%d\n", unsafe.Sizeof(nul))
	var nul1 = struct {
		A int
		B string
	}{}
	fmt.Printf("%d\n", unsafe.Sizeof(nul1))
	var pac = JSONPack{}
	fmt.Printf("%d\n", unsafe.Sizeof(pac))

	var i = uint64(1)
	fmt.Printf("uint64 %d\n", unsafe.Sizeof(i))
	var iPr = &i
	fmt.Printf("ptr %d\n", unsafe.Sizeof(iPr))
	var str = string("")
	fmt.Printf("str %d\n", unsafe.Sizeof(str))
	var lis = make([]string, 0, 1)
	fmt.Printf("%d\n", unsafe.Sizeof(lis))
	var ma = map[string]string{}
	fmt.Printf("%d\n", unsafe.Sizeof(ma))
	fmt.Printf("%d\n", 1^2)
}

func TestNewWSConn4(t *testing.T) {
	// a := 1.32
	// for i := 0; i< 100; i++ {
	// 	fmt.Println(int(a))
	// }
	for i := 0; i < 10; i++ {

		fmt.Println(uuid.New().String())

	}

}

func TestNewWSConn5(t *testing.T) {

	var y = float32(64)
	fmt.Println("????????????????????????", math.Sqrt(float64(y)))
	var x = y
	for (x*x - y) > 0.0001 { // ?????????????????????????????????
		x = (x + y/x) / 2
	}
	fmt.Println("?????????????????????", x)

	x2 := y
	i := math.Float32bits(y)
	i = 0x5f3759df - (i >> 1)
	x2 = math.Float32frombits(i)
	for a := 0; a < 3; a++ { // ?????????????????? ????????????
		x2 = x2 * (1.5 - (0.5 * y * x2 * x2))
	}
	fmt.Println("????????????3??????????????????", 1/x2)
}

func TestNewWSConn6(t *testing.T) {
	dp := make([]int, 100)
	dp[0] = 1
	var n2, n3, n5 = 0, 0, 0
	for i := 1; i < 100; i++ {
		dp[i] = int(math.Min(float64(2*dp[n2]), math.Min(float64(3*dp[n3]), float64(5*dp[n5]))))
		if dp[i] == 2*dp[n2] {
			n2++
		}
		if dp[i] == 3*dp[n3] {
			n3++
		}
		if dp[i] == 5*dp[n5] {
			n5++
		}
	}
	fmt.Println(dp[99])
}
