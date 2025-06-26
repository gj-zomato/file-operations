package file

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func PromptChoice(options []string) int {
	fmt.Println("\nWhat would you like to do?")
	for i, opt := range options {
		fmt.Printf("Press %d - %s\n", i+1, opt)
	}
	fmt.Print("Choice: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func CreateFile(filePath, fileName string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating the file:")
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%s file is created!\n", fileName)
	fmt.Printf("Path: %s", filePath)
	defer file.Close()
}

func MultiLineInput() string {
	var lines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading the file")
			fmt.Println(err)
			break
		}
		line = strings.TrimSpace(line)

		if line == "END" || line == "end" {
			break
		}
		lines = append(lines, line)
	}

	fullText := strings.Join(lines, "\n")
	fmt.Println("Content is added successfully!")
	return fullText
}

func WriteFile(filePath, input string) {
	err := os.WriteFile(filePath, []byte(input), 0644)
	if err != nil {
		fmt.Println("Error writing to file")
		fmt.Println(err)
		return
	}
}

func DisplayFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while displaying the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("> " + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading the file")
		fmt.Println(err)
	}
}

func UpdateFile(filePath, updatedContent string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while append")
		fmt.Println(err)
	}
	defer file.Close()

	if _, err := file.WriteString(updatedContent); err != nil {
		fmt.Println("Failed to write to file")
		fmt.Println(err)
	}
}

func ImportfromAPI(apiUrl, filePath string) error {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error while gettng the response!")
		fmt.Println()
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while opening the file!")
		fmt.Println(err)
	}
	defer file.Close()

	_, errors := io.Copy(file, resp.Body)
	if errors != nil {
		fmt.Println("Error while copying")
		fmt.Println(err)
	}
	return nil
}

func DeleteFile(filePath, toDelete string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error while deleting the file!")
		fmt.Println(err)
	}
}

func FormatSize(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	switch {
	case bytes >= GB:
		// switch case for bytes
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func ReadDir(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error while reading the directory")
		fmt.Println(err)
	}
	fmt.Println("\nHere are the files present in", filepath.Base(dirPath))
	for _, file := range files {
		if !file.IsDir() {
			fullPath := filepath.Join(dirPath, file.Name())
			info, err := os.Stat(fullPath)
			if err != nil {
				fmt.Println("Error while getting the size")
				fmt.Println(err)
			}
			fmt.Printf("- %s : %s\n", file.Name(), FormatSize(info.Size()))
		}
		if file.IsDir() {
			fullPath := filepath.Join(dirPath, file.Name())
			info, err := os.Stat(fullPath)
			if err != nil {
				fmt.Println("Error while getting the size")
				fmt.Println(err)
			}
			fmt.Printf("- %s : %s\n", file.Name(), FormatSize(info.Size()))
		}
	}
}
