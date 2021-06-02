package main

import "fmt"


func add(x , y int) int{
	return x + y
}

func calc(x int, y int) (int , int){
	add:= x + y;
	sub:= x - y;
	return add , sub
}

func calcMD(a , b int)(multi, divide int){
	multi = a * b;
	divide  = a / b;
	return 
}



func main(){

	var result = add(3, 4)
	var calcResult1 , calcResult2 = calc(10, 5);
	var calcMD1 , calcMD2 = calcMD(10, 5);

	fmt.Println("The Sum Of Given Two Numbers Are: ", result);
	fmt.Println("The calculation Of Given Two Numbers Are: ", calcResult1 , calcResult2);
	fmt.Println("The Multiple, Divide Of Given Two Numbers Are: ", calcMD1 , calcMD2);
	fmt.Println("Bye");
}
