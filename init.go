//go:build windows

package dll

func InitDllLoader() *DllLoader {
	dllLoader := DllLoader{
		Dlls:       make(map[string]map[int]*DllPackage),
		dllCounter: make(map[string]int),
	}
	return &dllLoader
}
