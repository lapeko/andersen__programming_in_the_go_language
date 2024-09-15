package main

import (
	"bufio"
	"fmt"
	"github.com/lapeko/andersen__programming_in_the_go_language/course1/lection11/app"
	"os"
	"strings"
)

func main() {
	taskApp := app.New()
	reader := bufio.NewReader(os.Stdin)
	var userInput string
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err, "Please try again")
			continue
		}

		userInput = strings.TrimSpace(line)
		strings.TrimSpace(userInput)

		if err != nil {
			fmt.Println(err, "Please try again")
			continue
		}

		if userInput == "Quit" {
			return
		}

		if userInput == "StartApp" {
			taskApp.StartApp(reader)
			return
		}

		fmt.Println("Unknown command. Please try again")
	}
}
