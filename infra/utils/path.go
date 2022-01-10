package utils

import (
	"path"
	"runtime"
)

func GetProjectAbPathByCaller() (abPath string) {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		filePath := path.Dir(filename)
		abPath = path.Dir(filePath)
	}
	return
}
