package main

import (
	"bufio"
	file "file-operations/function"
	check "file-operations/validation"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	invalidCounter := 0
	var dirPath string

	// Ask for directory path once at the beginning
	for {
		fmt.Print("📂 Enter your working directory path: ")
		fmt.Scan(&dirPath)

		if check.DirPathExist(dirPath) {
			fmt.Println("✅ Directory confirmed.")
			break
		} else {
			fmt.Println("❌ Invalid directory path. Please try again.")
		}
	}

	// Updated main menu with 'Change Directory'
	mainOptions := []string{
		"Create 🆕",
		"Edit ✏️",
		"View 📂",
		"Delete ❌",
		"Change Directory 🔄",
		"Quit 🚪",
	}
	createOptions := []string{"Import from API 🌐", "Create and Add Content 📝"}

	for {
		if invalidCounter >= 3 {
			fmt.Println("\n⚠️  You have entered invalid options 3 times. Exiting now!")
			break
		}

		fmt.Println("\n📋 Main Menu:")
		mainChoice := file.PromptChoice(mainOptions)

		switch mainChoice {

		case 1: // CREATE
			fmt.Println("\n📁 Create Options:")
			createChoice := file.PromptChoice(createOptions)

			switch createChoice {
			case 1: // Import from API
				var apiUrl string
				fmt.Print("🔗 Enter API URL: ")
				fmt.Scan(&apiUrl)

				fmt.Print("📝 Enter File Name (without extension): ")
				fileName, _ := reader.ReadString('\n')
				fileName = strings.TrimSpace(fileName)
				filePath := dirPath + "/" + fileName + ".json"

				if exists, err := check.FileExist(filePath); err != nil {
					fmt.Println("❗ Error:", err)
				} else if exists {
					fmt.Println("⚠️  File already exists.")
				} else {
					file.ImportfromAPI(apiUrl, filePath)
					fmt.Println("✅ API response stored in file.")
				}

			case 2: // Create and Add Content
				fmt.Print("📝 Enter File Name (without extension): ")
				fileName, _ := reader.ReadString('\n')
				fileName = strings.TrimSpace(fileName)
				filePath := dirPath + "/" + fileName + ".txt"

				if exists, err := check.FileExist(filePath); err != nil {
					fmt.Println("❗ Error:", err)
				} else if exists {
					fmt.Println("⚠️  File already exists.")
				} else {
					file.CreateFile(filePath, fileName)
					fmt.Println("✍️  Write your content (end input with a blank line):")
					content := file.MultiLineInput()
					file.WriteFile(filePath, content)
					fmt.Println("✅ File created and content added.")
				}

			default:
				fmt.Println("❌ Invalid Create option.")
				invalidCounter++
			}

		case 2: // EDIT
			file.ReadDir(dirPath)
			fmt.Print("✏️  Enter File Name to Edit (must end with .txt): ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			filePath := dirPath + "/" + fileName

			if !strings.HasSuffix(fileName, ".txt") {
				fmt.Println("🔒 Only .txt files can be edited.")
				break
			}

			if exists, err := check.FileExist(filePath); err != nil {
				fmt.Println("❗ Error:", err)
			} else if !exists {
				fmt.Println("⚠️  File does not exist.")
			} else {
				if check.IsEmpty(filePath) {
					fmt.Println("📄 File is empty. Enter new content:")
					updatedContent := file.MultiLineInput()
					file.UpdateFile(filePath, updatedContent)
				} else {
					fmt.Println("📄 Add Content (end input with a blank line):")
					file.DisplayFile(filePath)
					updatedContent := "\n" + file.MultiLineInput()
					file.UpdateFile(filePath, updatedContent)
				}
				fmt.Println("✅ File updated successfully.")
			}

		case 3: // VIEW
			file.ReadDir(dirPath)
			fmt.Print("🔍 Enter File Name to View: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			filePath := dirPath + "/" + fileName

			if exists, err := check.FileExist(filePath); err != nil {
				fmt.Println("❗ Error:", err)
				invalidCounter++
			} else if exists {
				fmt.Printf("\n📄 Content of %s:\n", fileName)
				file.DisplayFile(filePath)
			} else {
				fmt.Println("⚠️  File does not exist.")
				invalidCounter++
			}

		case 4: // DELETE
			file.ReadDir(dirPath)
			fmt.Print("🗑️  Enter File Name to Delete: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			filePath := dirPath + "/" + fileName

			if exists, err := check.FileExist(filePath); err != nil {
				fmt.Println("❗ Error:", err)
			} else if exists {
				file.DeleteFile(filePath, fileName)
				fmt.Println("✅ File deleted successfully.")
			} else {
				fmt.Println("⚠️  File does not exist.")
				invalidCounter++
			}

		case 5: // CHANGE DIRECTORY
			fmt.Print("🔄 Enter new working directory path: ")
			newPath, _ := reader.ReadString('\n')
			newPath = strings.TrimSpace(newPath)

			if check.DirPathExist(newPath) {
				dirPath = newPath
				fmt.Println("✅ Directory changed successfully.")
			} else {
				fmt.Println("❌ Invalid directory path.")
				invalidCounter++
			}

		case 6: // QUIT
			fmt.Println("👋 Thank you for using the File Manager. Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid option. Try again.")
			invalidCounter++
		}
	}
}
