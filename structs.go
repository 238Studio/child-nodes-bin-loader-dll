//go:build windows

package dll

import "syscall"

// DllPackage 一个dll包的结构体
type DllPackage struct {
	// 名字
	name string
	// id
	id int
	// 支持的函数名
	functions []string
	// 函数的入参类型 函数名-入参类型表
	functionsArgsTypes map[string][]string
	// 函数的返回值类型表 函数名-返回值类型表
	functionsReturnTypes map[string][]string
	// 其他信息
	info map[string]string
	// dll包
	dll *syscall.DLL
}

// DllLoader 加载dll包的结构体
type DllLoader struct {
	// 已加载的dll包 dllName -> id -> dll_packages
	Dlls map[string]map[int]*DllPackage
	// 包计数器 用于分配ID
	dllCounter map[string]int
}
