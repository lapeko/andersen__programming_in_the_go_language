package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		name string
		age  int
	)

	if _, err := fmt.Scan(&name, &age); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s has %d years old.\n")

	var text string

	if _, err := fmt.Fprintf(os.Stdin, text); err != nil {
		log.Fatalln(err)
	}

	log.Println("You input text is: ", text)
}
