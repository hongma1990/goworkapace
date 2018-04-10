package main

import (
	"fmt"
)

type def_func func(int) bool

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

func isEven(v int) bool {
	if v%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f def_func) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice is:", slice)
	fmt.Println("odd element is:", filter(slice, isOdd))
	fmt.Println("even element is:", filter(slice, isEven))
}
