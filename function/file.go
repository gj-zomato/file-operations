package file

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// CreateFile creates an empty file at the specified path
func CreateFile(filePath, fileName string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Failed to create the file:", err)
		return
	}
	defer file.Close()

	fmt.Println("\n✅ File has been created successfully!")
	fmt.Printf("📍 Path: %s\n", filePath)
}

// WriteFile writes the given content to the specified file
func WriteFile(filePath, input string) {
	err := os.WriteFile(filePath, []byte(input), 0644)
	if err != nil {
		fmt.Println("❌ Failed to write to file:", err)
	}
}

// DisplayFile prints the contents of a file line by line
func DisplayFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("❌ Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("> " + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error reading file content:", err)
	}
}

// UpdateFile appends new content to an existing file
func UpdateFile(filePath, updatedContent string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Failed to open file for appending:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(updatedContent); err != nil {
		fmt.Println("❌ Failed to write to file:", err)
	}
}

// ImportfromAPI fetches data from the provided API and stores it in a file
func ImportfromAPI(apiUrl, filePath string) error {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("❌ Failed to fetch from API:", err)
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Failed to create file for API response:", err)
		return err
	}
	defer file.Close()

	_, copyErr := io.Copy(file, resp.Body)
	if copyErr != nil {
		fmt.Println("❌ Failed to write API response to file:", copyErr)
		return copyErr
	}
	return nil
}

// DeleteFile removes the specified file
func DeleteFile(filePath, fileName string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("❌ Failed to delete file:", err)
	}
}

// FormatSize returns human-readable file size
func FormatSize(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

// ReadDir lists all files and folders in the given directory
func ReadDir(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("❌ Failed to read the directory:", err)
		return
	}

	fmt.Printf("\n📁 Contents of '%s':\n", filepath.Base(dirPath))
	for _, file := range files {
		fullPath := filepath.Join(dirPath, file.Name())
		info, err := os.Stat(fullPath)
		if err != nil {
			fmt.Println("⚠️  Could not get info for", file.Name())
			continue
		}
		fmt.Printf("• %s — %s\n", file.Name(), FormatSize(info.Size()))
	}
}
