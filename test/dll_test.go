package main

import (
	"testing"
	"unsafe"

	dll "github.com/238Studio/child-nodes-bin-loader-dll"
)

func TestName(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = uintptr(unsafe.Pointer(&a))
	app := dll.InitDllLoader()
	id, err := app.LoadBinPackage("./test")
	if err != nil {
		println(err.Error())
		return
	}
	re := make([]uintptr, 1)
	var str = "外部"
	println(&str)
	re[0] = (uintptr)(unsafe.Pointer(&str))

	binPackage := app.Dlls["test"][id]
	err = binPackage.Execute("Test1", args, uintptr(unsafe.Pointer(&re)))
	//	re := execute[0]
	println("mew")
	println(*(*string)(unsafe.Pointer(re[0])))

	app.ReleasePackage("test", id)
}
