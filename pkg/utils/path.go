package utils

import "os"

func PathExists(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true
		}
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
