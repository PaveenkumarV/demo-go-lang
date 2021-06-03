package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := [...]int{1, 2, 3, 4}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// arrays using String

	var names [3]string
	names[0] = "Paveen"
	names[1] = "Santhosh"
	names[2] = "5"

	fmt.Println(names)
	fmt.Printf("%v , %T\n", names[2], names[2])
	var num1, num2 = strconv.Atoi(names[2])
	if num2 != nil {
		return
	}
	fmt.Printf("%v, %T", num1, num1)
	fmt.Println(len(names))

	// twoDimentional Arrays

	var nestedArrays = [2][3]int{{0, 1, 2}, {2, 1, 0}}
	fmt.Println("nested Array ", nestedArrays)

	//copying Arrays

	var randomNum1 = [...]int{1, 4, 6}
	randomNum2 := randomNum1
	randomNum2[2] = 9
	fmt.Println("Original Copies ", randomNum1, randomNum2)

	//Coping Array By Referance

	var randomNums1 = [...]int{1, 4, 6}
	randomNums2 := &randomNums1 //copying by reference
	randomNums2[2] = 9
	fmt.Println("Copying By Reference Values ", randomNums1, randomNums2)

	sliceFunction()

}

func sliceFunction() {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num[6] = 18
	fmt.Println("From Slice Funtion: ", num[:])
	fmt.Println("From Slice Funtion: ", num[2:5])
	fmt.Println("From Slice Funtion: ", num[:6])
	fmt.Println("From Slice Funtion: ", num[5:])

	//Slice works with reference value while coping

	var randomNum1 = []int{1, 4, 6}
	randomNum2 := randomNum1
	randomNum2[2] = 9
	fmt.Println("Slice Copies ", randomNum1, randomNum2, "Capacity ", cap(randomNum2))

	// Using Make Funtion With Slice

	var makeWithSlice = make([]int, 3, 100) //second arg = size, third arg = capacity
	fmt.Println("Capacity ", cap(makeWithSlice), "length ", len(makeWithSlice))

	// why make shines and how append works

	emptyArray := []int{}
	fmt.Println(emptyArray)
	fmt.Printf("Capacity: %v\n", cap(emptyArray))
	fmt.Printf("Length: %v\n", len(emptyArray))

	emptyArray = append(emptyArray, 1)
	fmt.Println(emptyArray)
	fmt.Printf("Capacity: %v\n", cap(emptyArray))
	fmt.Printf("Length: %v\n", len(emptyArray))

	emptyArray = append(emptyArray, []int{2, 3, 4, 5}...) //Here we are using Spread Operater Like in Javascript
	fmt.Println(emptyArray)
	fmt.Printf("Capacity: %v\n", cap(emptyArray))
	fmt.Printf("Length: %v\n", len(emptyArray))

}
