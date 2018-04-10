package main

import (
	"fmt"
)

func main() {
	s := "kevin"
	b := []byte(s)
	b[0] = 'j'
	ns := string(b)
	fmt.Printf(ns)
}
