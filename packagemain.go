package main

import "fmt" // formatating text..

const (
	bmw = 1 << iota
	audi
	benz
	landrover
	toyota
)

func main() {
	var num int64 = bmw | landrover | toyota
	fmt.Printf("%b\n", num)
	if (num & landrover) == landrover {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
