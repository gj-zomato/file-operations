package main

import (
	"bufio"
	file "file-manager/function"
	check "file-manager/validation"
	"fmt"
	"os"
	"strings"
)

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)
	var api int

	functionalities := []string{"Import from API", "Create from Scratch", "View files in directory", "Delete"}
	vude := []string{"View", "Update", "Delete", "End"}

	api = file.PromptChoice(functionalities)
	if api == 1 {
		var apiUrl, dirPath string
		fmt.Println("Please provide the api url!")
		fmt.Printf("API Url: ")
		fmt.Scan(&apiUrl)

		fmt.Println("Please specify the directory path")
		fmt.Printf("Enter Path: ")
		fmt.Scan(&dirPath)

		fmt.Printf("Enter File Name: ")
		fileName, _ := reader.ReadString('\n')
		fileName = strings.TrimSpace(fileName)

		filePath := dirPath + "/" + fileName + ".json"

		exists, err := check.FileExist(filePath)
		if err != nil {
			fmt.Println("Error while checking")
			fmt.Println(err)
			return
		}
		if exists {
			fmt.Println("\nFile already exist")
		} else {
			file.ImportfromAPI(apiUrl, filePath)
			fmt.Println("API response stored!")
		}
	} else if api == 2 {
		var dirPath string
		fmt.Println("Please specify the directory path")
		fmt.Printf("Enter Path: ")

		fmt.Println("What would you like to name the file?")
		fmt.Printf("Enter File Name: ")
		fileName, _ := reader.ReadString('\n')
		fileName = strings.TrimSpace(fileName)
		filePath := dirPath + "/" + fileName

		exists, err := check.FileExist(filePath)
		if err != nil {
			fmt.Println("Error while checking")
			fmt.Println(err)
			return
		}
		if exists {
			fmt.Println("\nFile already exist")
			choice = file.PromptChoice(vude)
		} else {
			fmt.Println("\nCreating the file")
			file.CreateFile(filePath, fileName)
			fmt.Println("\nPlease write the content of the file")
			fullText := file.MultiLineInput()
			file.WriteFile(filePath, fullText)
		}

		choice = file.PromptChoice(vude)
		for {
			invalidCounter := 0
			if choice == 1 {
				fmt.Println("\nFile Content:")
				file.DisplayFile(filePath)
				break
			} else if choice == 2 {
				isEmpty := check.IsEmpty(filePath)
				if isEmpty {
					file.DisplayFile(filePath)
					updatedContent := file.MultiLineInput()
					file.UpdateFile(filePath, updatedContent)
				} else {
					file.DisplayFile(filePath)
					updatedContent := "\n" + file.MultiLineInput()
					file.UpdateFile(filePath, updatedContent)
				}
				choice = file.PromptChoice(vude)
			} else if choice == 3 {
				file.DeleteFile(filePath, fileName)
				break
			} else if choice == 4 {
				fmt.Println("Sure!")
				break
			} else {
				invalidCounter++
				if invalidCounter > 2 {
					fmt.Println("\nYou have entered the invalid choice 3 times!")
					break
				} else {
					fmt.Printf("\nDid not find what you were looking for!")
					fmt.Printf("\nPlease re-enter your choice: ")
					fmt.Scan(&choice)
				}
			}
		}

	} else if api == 3 {
		var dirPath string
		fmt.Println("Please enter the directory path")
		fmt.Printf("Enter path: ")
		fmt.Scan(&dirPath)
		file.ReadDir(dirPath)
	} else if api == 4 {
		var dirPath string
		fmt.Println("Enter the directory path of the file: ")
		fmt.Printf("Enter path: ")
		fmt.Scan(&dirPath)
		fmt.Println("Please enter the file you want to delete")
		fmt.Printf("Enter the file: ")
		toDelete, _ := reader.ReadString('\n')

		filePath := dirPath + "/" + toDelete
		filePath = strings.TrimSpace(filePath)
		exists, err := check.FileExist(filePath)
		if err != nil {
			fmt.Println("Error while deleting")
			fmt.Println(err)
		}
		if exists {
			file.DeleteFile(filePath, toDelete)
		} else {
			fmt.Println("File not found!")
		}

	} else {
		// to be ha
	}
}
