package class

import "fmt"

type Human struct {
	name string
}

func NewHuman(name string) *Human {
	return &Human{name: name}
}

func (human Human) GetName() string {
	return human.name
}

func (human *Human) SetName(name string) {
	human.name = name
}

func (human Human) Check() {
	fmt.Printf("%s is a Human.\n", human.name)
}

func (human Human) Speak() {
	fmt.Println("I am Human.")
}
