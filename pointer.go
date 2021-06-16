package main

import (
	"fmt"
)

func main() {
	a := 43
	// var b *int = &a
	var b = &a
	*b = 9
	fmt.Println(a)
	fmt.Println(*b)

	pointerWithStruct()
}

type details struct {
	name string
}

func pointerWithStruct() {
	// var d *details = &details{}
	var d *details
	// d = &details{"Paveen"}
	d = new(details)
	(*d).name = "PaveenV"
	fmt.Println((*d).name)
	// fmt.Println(*d)
}
