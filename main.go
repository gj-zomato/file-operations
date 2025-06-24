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

	reader := bufio.NewReader(os.Stdin)
	invalidCounter := 0

	functionalities := []string{"Import from API", "Create from Scratch", "View files in directory", "Delete a file", "Quit"}
	vude := []string{"View", "Update", "Delete", "Back to Main Menu"}

	for {
		if invalidCounter >= 3 {
			fmt.Println("\nYou have entered invalid options 3 times. Exiting now!")
			break
		}

		api := file.PromptChoice(functionalities)

		switch api {
		case 1: // Import from API
			var apiUrl, dirPath string
			fmt.Print("API Url: ")
			fmt.Scan(&apiUrl)

			fmt.Print("Enter Directory Path: ")
			fmt.Scan(&dirPath)

			if check.DirPathExist(dirPath) {
				fmt.Print("Enter File Name: ")
				fileName, _ := reader.ReadString('\n')
				fileName = strings.TrimSpace(fileName)
				filePath := dirPath + "/" + fileName + ".json"

				if exists, err := check.FileExist(filePath); err != nil {
					fmt.Println("Error:", err)
					continue
				} else if exists {
					fmt.Println("File already exists.")
				} else {
					file.ImportfromAPI(apiUrl, filePath)
					fmt.Println("API response stored.")
				}
			} else {
				fmt.Println("Invalid path.")
				invalidCounter++
			}

		case 2: // Create from Scratch
			var dirPath string
			fmt.Print("Enter Directory Path: ")
			fmt.Scan(&dirPath)

			fmt.Print("Enter File Name: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			filePath := dirPath + "/" + fileName

			if exists, err := check.FileExist(filePath); err != nil {
				fmt.Println("Error:", err)
				continue
			} else if exists {
				fmt.Println("File already exists.")
			} else {
				file.CreateFile(filePath, fileName)
				fmt.Println("\nWrite your content (end input with a blank line):")
				fullText := file.MultiLineInput()
				file.WriteFile(filePath, fullText)
			}

			for {
				choice := file.PromptChoice(vude)
			fileLoop:
				switch choice {
				case 1:
					file.DisplayFile(filePath)
				case 2:
					if check.IsEmpty(filePath) {
						fmt.Println("File is empty.")
						updatedContent := file.MultiLineInput()
						file.UpdateFile(filePath, updatedContent)
					} else {
						file.DisplayFile(filePath)
						updatedContent := "\n" + file.MultiLineInput()
						file.UpdateFile(filePath, updatedContent)
					}
				case 3:
					file.DeleteFile(filePath, fileName)
					break fileLoop
				case 4:
					fmt.Println("Returning to Main Menu.")
					break fileLoop
				default:
					fmt.Println("Invalid choice.")
					invalidCounter++
					if invalidCounter >= 3 {
						fmt.Println("Too many invalid attempts.")
						return
					}
				}
				if choice == 3 || choice == 4 {
					break
				}
			}

		case 3: // View Directory
			var dirPath string
			fmt.Print("Enter Directory Path: ")
			fmt.Scan(&dirPath)
			file.ReadDir(dirPath)

		case 4: // Delete File
			var dirPath string
			fmt.Print("Enter Directory Path: ")
			fmt.Scan(&dirPath)
			fmt.Print("Enter File Name to Delete: ")
			toDelete, _ := reader.ReadString('\n')
			filePath := strings.TrimSpace(dirPath + "/" + toDelete)

			if exists, err := check.FileExist(filePath); err != nil {
				fmt.Println("Error:", err)
			} else if exists {
				file.DeleteFile(filePath, toDelete)
			} else {
				fmt.Println("File does not exist.")
				invalidCounter++
			}

		case 5: // Quit
			fmt.Println("Thank you for using the File Manager. Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Try again.")
			invalidCounter++
		}
	}
}
