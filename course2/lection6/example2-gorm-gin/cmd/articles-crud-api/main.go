package main

import "github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/api"

func main() {
	a := api.New()
	a.Init()
	defer a.Storage.Close()
}
