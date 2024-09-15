package app

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func getFullDateString(dateString string) (string, error) {
	r := regexp.MustCompile(`^\d{1,4}-\d{1,2}-\d{1,2}$`)
	if !r.MatchString(dateString) {
		return "", errors.New("incorrect format")
	}

	dateParts := strings.Split(dateString, "-")

	return fmt.Sprintf(
		"%s-%s-%s",
		fillWithZeros(dateParts[0], 4),
		fillWithZeros(dateParts[1], 2),
		fillWithZeros(dateParts[2], 2),
	), nil
}

func fillWithZeros(srcString string, totalSize int) string {
	paddingSize := totalSize - len(srcString)
	return fmt.Sprintf("%s%s", strings.Repeat("0", paddingSize), srcString)
}

func deleteAllTasksForDate(taskMap map[string][]string, p *payload) int {
	tasks, ok := taskMap[p.date]
	if !ok {
		return 0
	}
	size := len(tasks)
	delete(taskMap, p.date)
	return size
}

func deleteOneTasksForDate(taskMap map[string][]string, p *payload) bool {
	tasks, ok := taskMap[p.date]
	if !ok {
		return false
	}
	indexToRemove := -1
	for idx, task := range tasks {
		if task != p.task {
			continue
		}
		indexToRemove = idx
		break
	}
	if indexToRemove == -1 {
		return false
	}
	if len(tasks) == 1 {
		delete(taskMap, p.date)
		return true
	}
	tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
	taskMap[p.date] = tasks
	return true
}
