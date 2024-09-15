package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// numbers
	num := 3
	fmt.Printf("Type of num is \"%T\" and it takes %d bytes\n", num, unsafe.Sizeof(num))

	// strings
	vitali := "Виталя"
	fmt.Printf("Length of word %s is %d bytes and has %d characters\n", vitali, len(vitali), utf8.RuneCountInString(vitali))

	for idx, letter := range vitali {
		fmt.Printf("Index: %d. Letter: %c\n", idx, letter)
	}

	var l rune = 'ł'
	fmt.Println(l)
	var l2 rune = 322
	fmt.Printf("Character %c has number code %d\n", l2, l2)
	fmt.Printf("%t", strings.Compare("Вася", "Грыша"))

	if num := 10; num%2 == 0 {
		fmt.Printf("%d is an even number", num)
	} else {
		fmt.Printf("%d is an odd number", num)
	}
}
