package cryptos

import (
	"crypto/md5"
	"fmt"
)

// Get32MD5 获取字符串32位md5
func Get32MD5(args string) string {
	h := md5.New()
	h.Size()
	h.Write([]byte(args))
	md5Str := fmt.Sprintf("%x", h.Sum(nil))
	return md5Str
}

// Get16MD5 获取字符串16位md5
func Get16MD5(args string) string {
	return Get32MD5(args)[8:24]
}
