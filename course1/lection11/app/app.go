package app

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type App struct {
	reader  *bufio.Reader
	taskMap map[string][]string
}

type payload struct {
	date string
	task string
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

			command, date := inputs[0], inputs[1]
			fullDate, err := getFullDateString(date)

			if err != nil {
				fmt.Println(inputErrorString)
				continue
			}

			p := &payload{date: fullDate}
			if len(inputs) > 2 {
				p.task = strings.Join(inputs[2:], " ")
			}

			switch command {
			case "Add":
				a.addTask(p)
			case "Del":
				a.delete(p)
			case "Find":
				a.findByDate(p)
			default:
				fmt.Println(inputs)
				fmt.Println(inputErrorString)
				continue
			}
		}
	}
}

func (a *App) addTask(p *payload) {
	if p.task == "" {
		fmt.Println(inputErrorString)
		return
	}

	tasks, ok := a.taskMap[p.date]

	if !ok {
		a.taskMap[p.date] = []string{p.task}
		return
	}

	tasks = append(tasks, p.task)
	sort.Strings(tasks)
	a.taskMap[p.date] = tasks
}

func (a *App) delete(p *payload) {
	if p.task == "" {
		deletedNum := deleteAllTasksForDate(a.taskMap, p)
		fmt.Printf("Deleted %d events\n", deletedNum)
		return
	}
	deleted := deleteOneTasksForDate(a.taskMap, p)
	if deleted {
		fmt.Println("Deleted successfully")
		return
	}
	fmt.Println("Event not found")
}

func (a *App) findByDate(p *payload) {
	tasks, ok := a.taskMap[p.date]
	if !ok {
		return
	}
	for _, task := range tasks {
		fmt.Println(task)
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
