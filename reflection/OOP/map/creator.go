package main

func Creator(animal string) Animal {
	var obj Animal
	obj = builder[animal]()
	return obj
}
