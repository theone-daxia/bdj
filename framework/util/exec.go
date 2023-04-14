package util

import "os"

// GetExecDirectory 获取当前执行程序的目录
func GetExecDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir + "/"
}
