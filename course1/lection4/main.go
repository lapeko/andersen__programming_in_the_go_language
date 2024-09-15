package main

import (
	"fmt"
)

func main() {
	// simple array
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	// computed size array
	arr2 := [...]int{1, 2}
	fmt.Println(arr == arr2)

	// slice
	arr3 := [...]int{3, 4, 5}
	slice := arr3[:]
	for idx, _ := range slice {
		slice[idx]++
	}
	fmt.Println(arr3)
	fmt.Printf("%v.\tSize: %d. Capacity: %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 7)
	fmt.Printf("%v.\tSize: %d. Capacity: %d\n", slice, len(slice), cap(slice))
}
