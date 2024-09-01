package main

import (
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/api"
	"log"
)

func main() {
	a := api.New()
	log.Fatal(a.Start())
}
