package main

import "fmt"

func main() {
	fallThrough()
}

func labeledLoop() {
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

func fallThrough() {
	var input int

	fmt.Println("Enter a number")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case input < 100:
		fmt.Println("Less then 100")
	case input < 200:
		fmt.Println("Less then 200")
		fallthrough
	case input >= 200:
		fmt.Println("Higher or equal 200")
	}
}
