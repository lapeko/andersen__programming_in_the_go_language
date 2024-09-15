package app

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type App struct {
	reader  *bufio.Reader
	taskMap map[string][]string
}

const inputErrorString = "Incorrect command or arguments. Please try again"

func New() *App {
	return &App{
		taskMap: map[string][]string{},
	}
}

func (a *App) StartApp(reader *bufio.Reader) {
	var userInput string
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
			a.printTasks()
		default:
			inputs := strings.Split(userInput, " ")

			if len(inputs) < 2 {
				fmt.Println(inputErrorString)
				continue
			}

			switch inputs[0] {
			case "Add":
				a.addTask(inputs)
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

func (a *App) printTasks() {
	var keys []string
	for key, _ := range a.taskMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, date := range keys {
		for _, task := range a.taskMap[date] {
			fmt.Println(date, task)
		}
	}
}

func (a *App) addTask(inputs []string) {
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

	val, ok := a.taskMap[dateInCorrectFormat]

	if !ok {
		a.taskMap[dateInCorrectFormat] = []string{task}
		return
	}

	val = append(val, task)
	sort.Strings(val)
}

//func delete(taskMap map[string][]string, inputs []string) {
//
//}
