package utils

import "os"

// Does the file exist
func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

// Is it a directory
func IsDir(name string) bool {
	if !FileExist(name) {
		return false
	}

	stat, _ := os.Stat(name)
	return stat.IsDir()
}

// Is it a file
func IsFile(name string) bool {
	if !FileExist(name) {
		return false
	}

	return !IsDir(name)
}
