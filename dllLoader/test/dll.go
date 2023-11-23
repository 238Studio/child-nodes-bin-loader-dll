package main

import "C"
import (
	"unsafe"
)

//export Release
func Release() {
	println("已释放")
}

//export Test0
func Test0() {
	println("测试")
}

// Test1 测试go传入传出 和内部操作指针
//
//export Test1
func Test1(re uintptr, args uintptr) uintptr {
	arg := (*[]uintptr)(unsafe.Pointer(args))
	str := *(*string)(unsafe.Pointer((*arg)[0]))
	b := str + "mew"
	r := (*[]uintptr)(unsafe.Pointer(re))
	println("内部")
	println(*(*string)(unsafe.Pointer((*r)[0])) + "mew")
	*(*string)(unsafe.Pointer((*r)[0])) = b
	return uintptr(unsafe.Pointer(&re))
}

//ex

func main() {

}
