package main

import (
	"fmt"
	"os"
)

func main() {
	var input1, input2, input3 string

	fmt.Println("Введите три слова: ")

	_, err := fmt.Fscanf(os.Stdin, "%s %s %s\n", &input1, &input2, &input3)

	if err != nil {
		fmt.Println(err)
		return
	}

	if (input1 == "один" || input1 == "раз") && input2 == "два" && input3 == "три" {
		fmt.Println("Ok")
		return
	}
	if input1 == "1" && input2 == "2" && input3 == "3" {
		fmt.Println("Ok")
		return
	}
	fmt.Println("Not Ok")
}
