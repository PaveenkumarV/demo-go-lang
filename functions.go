package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	Operator := "+"
	description, result := calcu(num1, num2, Operator)
	fmt.Println(description, result)
}

func add(x, y float32) (string, float32) {
	return "Addition Of: ", x + y
}
func sub(x, y float32) (string, float32) {
	return "Substration Of: ", x - y
}
func multi(x, y float32) (string, float32) {
	return "Multiplication Of: ", x * y
}
func divide(x, y float32) (string, float32) {
	return "Division Of ", x / y
}

func calcu(x, y int, TyOfOp string) (string, float32) {

	xi := float32(x)
	yi := float32(y)

	var result float32
	var description string

	switch TyOfOp {
	case "+":
		description, result = add(xi, yi)
	case "-":
		description, result = sub(xi, yi)
	case "*":
		description, result = multi(xi, yi)
	case "/":
		description, result = divide(xi, yi)
	}
	return description, result
}
