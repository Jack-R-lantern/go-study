package main

func Creator(animal string) Animal {
	var obj Animal
	switch animal {
	case "Dog":
		obj = CreateDog()
	case "Cat":
		obj = CreateCat()
	case "Cow":
		obj = CreateCow()
		//case "Horse":
		//	obj = CreateHorse()
	}
	return obj
}
