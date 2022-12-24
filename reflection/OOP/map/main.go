package main

import "os"

type createAnimal func() Animal

var builder map[string]createAnimal

func init() {
	builder = make(map[string]createAnimal)

	builder["Dog"] = CreateDog
	builder["Cat"] = CreateCat
	builder["Cow"] = CreateCow
}

func main() {
	Creator(os.Args[1]).Barking()
}
