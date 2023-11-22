package stringx

import "unsafe"

// 不安全地将string类型转换为byte类型
func UnsafeStringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// 不安全地将byte类型转换为string类型
func UnsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 安全地将string类型转换为byte类型
func SafeStringToBytes(s string) []byte {
	return []byte(s)
}

// 安全地将byte类型转换为string类型
func SafeBytesToString(b []byte) string {
	return string(b)
}
