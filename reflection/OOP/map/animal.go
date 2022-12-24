package main

import "fmt"

type Animal interface {
	Barking()
}

type Dog struct {
}

func (d *Dog) Barking() {
	fmt.Printf("왈왈")
}

type Cat struct {
}

func (c *Cat) Barking() {
	fmt.Printf("야옹")
}

type Cow struct {
}

func (c *Cow) Barking() {
	fmt.Printf("음메")
}

type Horse struct {
}

func (h *Horse) Barking() {
	fmt.Printf("휘위잉")
}
