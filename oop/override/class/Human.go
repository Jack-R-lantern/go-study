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

type Man struct {
	*Human
}

func NewMan(name string) *Man {
	return &Man{Human: NewHuman(name)}
}

func (man Man) Check() {
	fmt.Printf("%s is a Man.\n", man.name)
}
