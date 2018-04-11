package main

import (
	"fmt"
)

func double(a *int) {
	*a += *a
	a = nil
}

func main() {
	a := 3
	fmt.Println("pointer:", &a)
	double(&a)
	fmt.Println("value:", a)

	fmt.Println("pointer:", &a == nil)
}
