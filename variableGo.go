package main

import "fmt"

func main() {

	//  num1 := 10  // short hand declaration.
	// // var num2 = 20
	// var num2 int
	// num2 = 20

	var num1, num2 int

	num1 = 20
	num2 = 20

	var num3 complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(num3), real(num3))
	fmt.Printf("%v, %T\n", imag(num3), imag(num3))

	fmt.Println(num1 + num2)
}
