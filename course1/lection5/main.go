package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	testLine := "Тестовая строка"

	for i := 0; i < len(testLine); i++ {
		fmt.Printf("%c", testLine[i])
	}
	fmt.Println()

	runeSlice := []rune(testLine)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c", runeSlice[i])
	}
	fmt.Println()

	fmt.Printf("string size of testLine is %d\n", len(testLine))
	fmt.Printf("rune size of testLine is %d\n", utf8.RuneCountInString(testLine))

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter text")
		input.Scan()
		line := input.Text()
		if line == "Break" {
			break
		}
		fmt.Println(line)
	}
}
