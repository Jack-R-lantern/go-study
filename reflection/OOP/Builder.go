package main

import (
	"fmt"
	"reflect"
)

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

func Creator(animal string, builder *Builder) Animal {
	var obj Animal
	rt := reflect.ValueOf(builder)

	m := rt.MethodByName(fmt.Sprintf("Create%s", animal))
	obj = m.Call([]reflect.Value{})[0].Interface().(Animal)
	return obj
}
