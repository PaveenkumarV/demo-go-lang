package main

import "fmt"

func main() {
	sampleNum := 4
	switch sampleNum {
	case 1, -1:
		fmt.Println("The Given Number Is - Or + 1")
		break
	case 2:
		fmt.Println("The Given Number Is ", sampleNum)
		break
	case 3:
		{
			fmt.Println("The Given Number Is ", sampleNum)
		}
		break
	case 4:
		fmt.Println("The Given Number Is ", sampleNum)
		break
	case 5:
		fmt.Println("The Given Number Is ", sampleNum)
		break

	default:
		{
			fmt.Println("Switch Case Supported For Values That are Lesser Than 5")
		}
	}
}
