package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("Enter a word")

	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		log.Fatalln(err)
	}

	runes := []rune(strings.ToLower(input))

	if len(runes) < 2 {
		fmt.Println("НЕ СОГЛАСЕН")
		return
	}

	first, second := runes[0], runes[len(runes)-1]

	if first != 'д' || second != 'а' {
		fmt.Println("НЕ СОГЛАСЕН")
		return
	}

	fmt.Println("СОГЛАСЕН")
}
