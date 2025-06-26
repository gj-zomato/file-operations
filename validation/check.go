package check

import (
	"errors"
	"fmt"
	"os"
)

// IsEmpty checks if the file at the given path is empty
func IsEmpty(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("❌ Failed to check file status:")
		fmt.Println(err)
		return false
	}
	return info.Size() == 0
}

// FileExist returns true if the file exists at the given path
func FileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// DirPathExist checks if the provided path is a valid directory
func DirPathExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Println("❌ Error accessing the directory:")
		fmt.Println(err)
		return false
	}
	if !info.IsDir() {
		fmt.Println("⚠️  The given path is not a directory.")
		return false
	}
	return true
}

// APIresponse – Placeholder for future implementation
func APIresponse(apiUrl string) {
	fmt.Println("🚧 API response check not implemented yet.")
}
