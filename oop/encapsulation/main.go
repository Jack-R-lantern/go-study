package main

import (
	"encapsulation/class"
	"fmt"
)

func main() {
	human := class.NewHuman("alice")
	fmt.Println(human.GetName())
	human.SetName("bob")
	fmt.Println(human.GetName())
}
