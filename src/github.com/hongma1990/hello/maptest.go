package main

import (
	"fmt"
	// "map"
)

func main() {
	type T struct {
		age string
	}
	ma := map[string]T{}
	ma["John"] = T{age: "23"}

	ml := map[string]string{
		"1": "one",
	}

	fmt.Println(ma)
	fmt.Println(ml)

	ma["John"] = T{age: "96"}

	fmt.Println(ma)
}
