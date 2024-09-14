package main

import "fmt"

func main() {
	var input string

	for {
		_, err := fmt.Scanf("%s\n", &input)

		if err != nil {
			fmt.Println(err)
			return
		}

		if input == "0" {
			return
		}

		fmt.Println(input)
	}
}
