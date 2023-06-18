package utils

import "os"

// FileExists checks if a file exists at the given path
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// CreateDirectory creates a directory at the given path
func CreateDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
