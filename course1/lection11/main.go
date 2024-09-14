package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const inputErrorString = "Incorrect command or arguments. Please try again"

func main() {
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
			startApp(reader)
			return
		}

		fmt.Println("Unknown command. Please try again")
	}
}

func startApp(reader *bufio.Reader) {
	var userInput string
	taskMap := map[string][]string{}
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err, "Please try again")
			continue
		}

		userInput = strings.TrimSpace(line)
		strings.TrimSpace(userInput)

		if err != nil {
			log.Println(err, "Please try again")
			continue
		}

		if userInput == "Quit" {
			return
		}

		if userInput == "Print" {
			printTasks(taskMap)
		}

		inputs := strings.Split(userInput, " ")

		if len(inputs) < 2 {
			fmt.Println(inputErrorString)
			continue
		}

		switch inputs[0] {
		case "Add":
			addTask(taskMap, inputs)
		case "Del":
			break
		case "Find":
			break
		default:
			fmt.Println(inputErrorString)
			continue
		}
	}
}

func printTasks(taskMap map[string][]string) {
	var keys []string
	for key, _ := range taskMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, date := range keys {
		for _, task := range taskMap {
			fmt.Println(date, task)
		}
	}
}

func addTask(taskMap map[string][]string, inputs []string) {
	if len(inputs) < 3 {
		fmt.Println(inputErrorString)
		return
	}

	val, ok := taskMap[inputs[1]]

	if !ok {
		taskMap[inputs[1]] = []string{inputs[2]}
		return
	}

	val = append(val, inputs[2])
	sort.Strings(val)
}
