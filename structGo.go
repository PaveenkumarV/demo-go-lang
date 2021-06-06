package main

import (
	"fmt"
	"reflect"
)

// Basic Struct

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

	compositeStruct() //Function Call

	anonymousStruct() // anonymousStruct Function Call
}

//creating AnonymousStruct

func anonymousStruct() {
	name := struct {
		firstName string
		lastName  string
	}{firstName: "Paveen", lastName: "Kumar"} // anonymous struct
	fmt.Println(name.firstName)
}

// Like Inheritance, In Struct.

type India struct {
	food          string
	passPortNum   int
	primeMinister string
}

type America struct {
	India
	president string
}

func compositeStruct() {

	country := America{
		president: "Biden",
		India:     India{passPortNum: 42324343, food: "Pizza"},
	}
	fmt.Println(country)

	retunedFields := usingTagsInStruct() //Function Call
	fmt.Println(retunedFields)
}

// Struct Using Tags

func usingTagsInStruct() string {
	type Details struct {
		name        string `required max:"100"`
		nationality string
	}
	details := reflect.TypeOf(Details{})
	// fmt.Println(details)
	field, _ := details.FieldByName("name")
	return string(field.Tag)
}
