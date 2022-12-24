package main

import (
	"fmt"
	"log"
	"reflect"
)

func Creator(animal string, builder *Builder) Animal {
	var obj Animal
	rt := reflect.ValueOf(builder)

	m := rt.MethodByName(fmt.Sprintf("Create%s", animal))
	if reflect.ValueOf(m).IsNil() {
		log.Fatal("invalid type")
	}
	obj = m.Call([]reflect.Value{})[0].Interface().(Animal)

	return obj
}
