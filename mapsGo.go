package main

import "fmt"

func main() {

	likeObjects := map[string]int{
		"Paveen":   9121999,
		"Goutham":  4101999,
		"Harish":   17032000,
		"Santhosh": 15032003,
	}
	fmt.Println(likeObjects["Paveen"])
	fmt.Println(likeObjects["India"]) // Returns 0
	_, ok := likeObjects["India"]
	fmt.Println(ok)

	mapUsingMake()
}

func mapUsingMake() {
	likeObjects := map[int]string{}
	likeObjects = map[int]string{
		1: "Paveen",
		2: "Harish",
		3: "Goutham",
		4: "Santhosh",
		5: "GouVanth",
	}

	copies := likeObjects //Works with reference
	delete(likeObjects, 5)

	fmt.Println("From mapUsingMake Function: ", likeObjects[1])
	fmt.Println("From mapUsingMake Function, Length: ", len(likeObjects))
	fmt.Println("From mapUsingMake Function: ", likeObjects)
	fmt.Println("From mapUsingMake Function: ", copies)
}
