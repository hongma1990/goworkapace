package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	Skills
	grade  int
	phone  string
	parent map[string]string
}

type Worker struct {
	Human
	company string
}

func (h *Human) EatWhat() {
	fmt.Printf("I am %s,eat %s\n", h.name, "human")
}

func (w *Worker) EatWhat() {
	fmt.Printf("I am %s, eat %s\n", w.name, "worker")
}

func main() {
	kevin := Student{Human: Human{name: "kevin", age: 18}, grade: 4, phone: "13800138000"}
	var kevin2 Student
	kevin2.name = "kevin2"
	fmt.Println("kevin2:", kevin2)
	fmt.Println("kevin.name:", kevin.name)
	fmt.Println("kevin.Human.name:", kevin.Human.name)
	kevin.Human.phone = "10086"
	kevin.phone = "10065"
	fmt.Println("kevin.Human.phone:", kevin.Human.phone)
	fmt.Println("kevin.phone:", kevin.phone)
	kevin.Skills = []string{"play basketball"}
	kevin.Skills = append(kevin.Skills, "play tennis", "coding Java")
	fmt.Println("lenth of Skills:", len(kevin.Skills))
	kevin.parent = make(map[string]string)
	kevin.parent["mother"] = "mymother"
	kevin.parent["father"] = "myfather"
	fmt.Println("kevin.parent.lenth:", len(kevin.parent))
	fmt.Println("kevin.father:", kevin.parent["father"])
	fmt.Println(kevin)
	worker := Worker{Human: Human{name: "kevin_worker", age: 28}, company: "Netease Inc"}

	kevin.EatWhat()
	worker.EatWhat()
}
