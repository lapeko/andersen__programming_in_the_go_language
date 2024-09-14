package main

import "fmt"

func main() {
	// simple array
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1)

	// computed size array
	arr2 := [...]int{1, 2}
	fmt.Println(arr1 == arr2)
}
