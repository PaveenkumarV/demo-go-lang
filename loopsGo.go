package main
import "fmt"
func main(){

	

	for i := 1; i<=5 ; i++{
		var temp = "th"
		
		switch i{
		case 1:
			temp = "st"
			break
		case 2:
			temp = "nd"
			break
		case 3:
			temp = "rd"
		}
		
		fmt.Println(i, temp + " Time")
	}
}