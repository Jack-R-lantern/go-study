package main

func CreateDog() Animal {
	return &Dog{}
}

func CreateCat() Animal {
	return &Cat{}
}

func CreateCow() Animal {
	return &Cow{}
}

func CreateHorse() Animal {
	return &Horse{}
}
