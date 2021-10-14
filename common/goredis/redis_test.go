package goredis

import (
	"context"
	"fmt"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	client, err := NewRedisClient(context.Background())
	if err != nil {
		return
	}
	fmt.Println(client.Keys(context.Background(), "data_pay_*").Result())
	//conn, err := NewConnObj(context.Background())
	//if err != nil {
	//	return
	//}
	//defer conn.conn.Close()
	////err = conn.HMSet("dict", map[string]interface{}{"111": 100})
	////if err != nil {
	////	fmt.Println(err)
	////	return
	////}

	//fmt.Println(reflect.TypeOf(*a).Field(0).Name)
	//fmt.Println(reflect.ValueOf(a))
	// fmt.Println(conn.SetExpire("dict", 10 * time.Second))

	// conn.Set(context.Background(), "str", "h", 10*time.Second)

	// 数组
	//b := conn.RPush(context.Background(), "11111", "helloword")
	//b = conn.RPush(context.Background(), "11111", "helloword2")
	//fmt.Println(b.Result())
	//fmt.Println(conn.LRange(context.Background(), "11111", 0, 1).Result())

	// 字典
	//conn.HMSet(context.Background(), "dict", map[string]interface{}{"age": 10, "name": "word"})
	//fmt.Println(conn.HMGet(context.Background(), "dict", "name").Result())
	//fmt.Println(conn.HGetAll(context.Background(), "dict").Result())

	// 集合
	//conn.SAdd(context.Background(), "stringList", "11", "22", "33")
	//fmt.Println(conn.SMembers(context.Background(), "stringlist").Result())

	// 排序
	//conn.ZAdd(context.Background(), "zs", &redis.Z{Score: float64(10), Member: "a_3"}, &redis.Z{Score: float64(20), Member: "a_4"})
	//fmt.Println(conn.ZRank(context.Background(),"zs", "a_1").Result())
	//fmt.Println(conn.ZRevRange(context.Background(), "zs", 0, 2).Result())

}
