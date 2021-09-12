package config

import (
	"fmt"
	"golang_structure_project/router"
)

func Cobain(red, route string) {
	fmt.Println("Ini nyobain manggil " + red)
	r1 := router.NewRoute(route)
	fmt.Println(r1)
}
