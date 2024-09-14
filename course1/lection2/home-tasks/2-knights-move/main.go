package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int

	fmt.Println("Enter x1, y1, x2, y2")

	_, err := fmt.Scanf("%d\n%d\n%d\n%d\n", &x1, &y1, &x2, &y2)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	dX, dY := abs(x2-x1), abs(y2-y1)

	if (dX == 1 || dY == 1) && dX+dY == 3 {
		fmt.Println("Ok")
		return
	}

	fmt.Println("Not Ok")
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
