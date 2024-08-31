package main

import "fmt"

func main() {
	if num := 10; num%2 == 0 {
		fmt.Printf("%d is an even number", num)
	} else {
		fmt.Printf("%d is an odd number", num)
	}
}
