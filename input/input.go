package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptChoice shows a list of options and captures user input
func PromptChoice(options []string) int {
	fmt.Println("\n🤖 What would you like to do?")
	for i, opt := range options {
		fmt.Printf("Press %d - %s\n", i+1, opt)
	}
	fmt.Printf("👉 Enter your choice: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

// MultiLineInput captures multiline input from the user until "END"
func MultiLineInput() string {
	var lines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			break
		}
		line = strings.TrimSpace(line)
		if strings.EqualFold(line, "end") {
			break
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
