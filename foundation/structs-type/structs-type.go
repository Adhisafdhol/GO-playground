package main

import (
	"fmt"
)

func main() {
	// A struct type is defined with the keyword type
	type person struct {
		name        string
		age         int
		dateOfBirth string
	}

	var anthony person
	bonny := person{}
	fmt.Println("anthony:", anthony)
	fmt.Println("bonny:", bonny)
	// Define struct
	// All field must be specified
	celine := person{
		"Celine",
		36,
		"12/07/1989",
	}
	fmt.Println("celine:", celine)
	// Defined struct
	// You can specify the fields in any order
	// You can left out some field and it'll be set to its zero value
	demitry := person{
		name:        "Demitry",
		age:         20,
		dateOfBirth: "21/08/2005",
	}
	fmt.Println("demitry:", demitry)
	// Access and write a struct
	demitry.name = "Demitry Johnson"
	fmt.Println(demitry.name)
	// Anynomous structs
	// This create a single instance of the struct
	// Used for translating external data into a struct or vice versa
	var anynomousStruct struct {
		name string
		hobby string
	}
	
	anynomousStruct.name = "Adrian"
	anynomousStruct.hobby = "Swimming"
	
	anynomousStructTwo := struct {
			name string
			hobby string
	}{
		name: "Adrian",
		hobby: "Running",
	}
	fmt.Println("anynomousStructTwo:", anynomousStructTwo)

	// Comparing structs
	// If all strucks' fields are comparable, the struct is comparable 
	fmt.Println(demitry == celine)
	uncomparableStruct := struct{
		world map[string]int
	}{
		world: map[string]int{
		"one": 1},
	}
	uncomparableStructTwo := struct{
		world map[string]int
	}{
		world: map[string]int{
		"one": 1},
	}
	fmt.Println(uncomparableStruct, uncomparableStructTwo)
	// Otherwise it cannot be compared
	// fmt.Println(uncomparableStructTwo == uncomparableStruct)\
	// Cannot compare two different struct types
	// fmt.Println(anynomousStructTwo == anthony)

	// You can use type conversion on struct if the fields' names match
	// in the right order and have no additional field

	type personTwo struct {
		name string
		age int
		dateOfBirth string
	}

	erin := personTwo{
		name: "erin",
		age: 19,
		dateOfBirth: "21/12/2005",
	}
	fmt.Println(personTwo(celine) == erin )
	/* If one of the structs to be compared is anonymous you don't need
	 type conversion if all fields' names, order match and no additional
	fields */
	type structTwo struct {
		name string
		hobby string
	}

	adrian := structTwo{
		name: "Adrian",
		hobby: "Running",
	}
	fmt.Println(anynomousStructTwo == adrian)
}
