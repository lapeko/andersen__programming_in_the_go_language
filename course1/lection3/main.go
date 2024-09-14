package main

import "fmt"

func main() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > i {
				break
			}
			if i == 8 && j == 2 {
				break outer
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
