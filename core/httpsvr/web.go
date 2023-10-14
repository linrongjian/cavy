package httpsvr

import (
	"crypto/md5"
	"fmt"
	"strings"
	"unsafe"
)

// GenSign 生成签名
func (r *Request) genSign(vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	array := append(vals, SignKey)
	str := strings.Join(array, "+")
	hash := md5.Sum(r.stringToByte(str))
	return fmt.Sprintf("%x", hash)
}

// GenSign 生成签名
func GenSign(vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	array := append(vals, SignKey)
	str := strings.Join(array, "+")
	hash := md5.Sum(stringToByte(str))
	return fmt.Sprintf("%x", hash)
}

// 更高效率的byte转string
func (r *Request) byteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 更高效率的string转byte
func (r *Request) stringToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 更高效率的string转byte
func stringToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
