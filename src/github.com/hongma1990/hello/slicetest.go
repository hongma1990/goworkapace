package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}
	persons := []Person{
		{name: "Jack", age: 23},
		{name: "Rose", age: 25},
	}

	for _, p := range persons {
		persons[0].name = "kevin"
		fmt.Println(p.name, p.age)

	}
	fmt.Println(persons)
	persons = apepnd(persons, Person{name: "Lillian", age: 32})
	fmt.Println(persons)
}
