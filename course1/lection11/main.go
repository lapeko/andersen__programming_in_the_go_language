package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

		switch userInput {
		case "Quit":
			return
		case "Print":
			printTasks(taskMap)
		default:
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
}

func printTasks(taskMap map[string][]string) {
	var keys []string
	for key, _ := range taskMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, date := range keys {
		for _, task := range taskMap[date] {
			fmt.Println(date, task)
		}
	}
}

func addTask(taskMap map[string][]string, inputs []string) {
	if len(inputs) < 3 {
		fmt.Println(inputErrorString)
		return
	}

	r := regexp.MustCompile(`^\d{1,4}-\d{1,2}-\d{1,2}$`)
	date, task := inputs[1], inputs[2]

	if !r.MatchString(date) {
		log.Println("Provide a date in correct format")
		return
	}

	dateParts := strings.Split(date, "-")

	dateInCorrectFormat := fmt.Sprintf(
		"%s-%s-%s",
		fillWithZeros(dateParts[0], 4),
		fillWithZeros(dateParts[1], 2),
		fillWithZeros(dateParts[2], 2),
	)

	val, ok := taskMap[dateInCorrectFormat]

	if !ok {
		taskMap[dateInCorrectFormat] = []string{task}
		return
	}

	val = append(val, task)
	sort.Strings(val)
}

func fillWithZeros(srcString string, totalSize int) string {
	paddingSize := totalSize - len(srcString)
	return fmt.Sprintf("%s%s", strings.Repeat("0", paddingSize), srcString)
}
