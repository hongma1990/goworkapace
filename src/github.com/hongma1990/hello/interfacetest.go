package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) String() string {
	return h.name + " - " + strconv.Itoa(h.age) + " years - " + h.phone
}

func main() {
	list := make(List, 3)
	tom := Human{name: "tom", age: 23, phone: "1380013800"}
	list[0] = 1
	list[1] = "two"
	list[2] = tom
	fmt.Println("use if...")
	for _, elment := range list {
		if value, ok := elment.(int); ok {
			fmt.Println("I am type of int, value is:", value)
		} else if value, ok := elment.(string); ok {
			fmt.Println("I am type of string, value is:", value)
		} else if value, ok := elment.(Human); ok {
			fmt.Println("I am type of Human, value is:", value)
		}
	}
	fmt.Println("use switch")
	for index, elment := range list {
		switch value := elment.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is an Human and its value is %s\n", index, value.String())
		}
	}

}
