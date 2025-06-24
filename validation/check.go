package check

import (
	"errors"
	"fmt"
	"os"
)

func IsEmpty(filePath string) bool {
	empty, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error while checking!")
		fmt.Println(err)
	}
	return empty.Size() == 0
}

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

func DirPathExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Println("Error accessing path")
		fmt.Println(err)
		return false
	}
	fmt.Printf("✅ Valid path (%s)\n", func() string {
		if info.IsDir() {
			return "directory"
		}
		return "file"
	}())
	return true

}

func APIresponse(apiUrl string) {
	// to be done
}
