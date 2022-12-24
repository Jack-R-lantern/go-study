package main

type Builder struct {
}

func (b *Builder) CreateDog() Animal {
	return &Dog{}
}

func (b *Builder) CreateCat() Animal {
	return &Cat{}
}

func (b *Builder) CreateCow() Animal {
	return &Cow{}
}

func (b *Builder) CreateHorse() Animal {
	return &Horse{}
}
