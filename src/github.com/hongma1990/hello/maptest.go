package main

import (
	"fmt"
	// "map"
)

func main() {
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2

	fmt.Printf(len(numbers))

	value, ok := numbers["one1"]

	fmt.Printf(string(value) + "\n")
	fmt.Print(ok)
}
