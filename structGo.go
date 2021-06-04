package main

import "fmt"

type Details struct {
	companyName string
	location    string
	empDetails  map[string]string
}

func main() {
	details := Details{
		companyName: "Simplify",
		location:    "Bangalore",
		empDetails: map[string]string{
			"empName":       "Goutham",
			"empLocation":   "Bangalore",
			"empDepartment": "Developer",
		},
	}
	fmt.Println(details.empDetails["empName"])
	fmt.Println(details.companyName)

	compositeStruct()
}

type India struct {
	food          string
	passPortNum   int
	primeMinister string
}

type America struct {
	India
	pizza string
}

func compositeStruct() {

	country := America{
		pizza: "Cheese Pizza",
		India: India{passPortNum: 42324343, primeMinister: "Biden"},
	}
	fmt.Println(country)
}
