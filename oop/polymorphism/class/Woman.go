package class

import "fmt"

type Woman struct {
	name string
}

func NewWoman(name string) *Woman {
	return &Woman{name: name}
}

func (woman Woman) Speak() {
	fmt.Println("I am Woman.")
}
