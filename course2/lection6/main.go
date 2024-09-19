package main

func factorial(num int) int {
	res := 1
	if num < 2 {
		return res
	}
	for i := 2; i <= num; i++ {
		res *= i
	}
	return res
}
