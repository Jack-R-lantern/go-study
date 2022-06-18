package class

import "fmt"

type Man struct {
	*Human
}

func NewMan(name string) *Man {
	return &Man{Human: NewHuman(name)}
}

func (man Man) Check() {
	fmt.Printf("%s is a Man.\n", man.name)
}
