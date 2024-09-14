package main

import "fmt"

func main() {
	var userNum uint

	fmt.Println("Enter positive number")

	_, err := fmt.Scanf("%d\n", &userNum)

	if err != nil {
		fmt.Println("Inout error: ", err)
		return
	}

	for i, k := uint(0), userNum-1; i < userNum; i, k = i+1, k-1 {
		for j := uint(0); j < k; j++ {
			fmt.Print(" ")
		}

		current, prev := 1, 1
		for j := uint(0); j < i+1; j++ {
			fmt.Print(current, " ")
			if float32(j+1) == float32(i+1)/2 {
				continue
			}
			temp := current
			if j < i/2 {
				current += prev
				prev = temp
			} else {
				current = prev
				prev = temp - current
			}
		}

		fmt.Println()
	}
}
