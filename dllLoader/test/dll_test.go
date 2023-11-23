package main

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-hdex-loader/dllLoader"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	args := make([]uintptr, 1)
	a := "helloworld"
	args[0] = uintptr(unsafe.Pointer(&a))
	app := dllLoader.InitDllLoader()
	hexPackage, err := app.LoadHexPackage("C:\\Users\\real_common_cat\\Desktop\\childNodes\\child-nodes-hex-loader\\dllLoader\\test\\dll")
	if err != nil {
		println(err.Error())
		return
	}
	re := make([]uintptr, 1)
	var str = "外部"
	println(&str)
	re[0] = (uintptr)(unsafe.Pointer(&str))
	err = hexPackage.Execute("Test1", args, uintptr(unsafe.Pointer(&re)))
	//	re := execute[0]
	println("mew")
	println(*(*string)(unsafe.Pointer(re[0])))
}
