package main

import "fmt"

func main () {
	// Booleans
	var isTrue bool = false
	fmt.Println(isTrue)

	// Numberics
	// int8
	// int16
	// int32
	// int64
	// uint8
	// uint16
	// uint32
	// uint64
	
	// Special integer types
	/* int 32-bit signed integer in 32-bit cpu and 64-bit signed integer in
	 64-bit-cpu */
	// byte an alies for uint8
	
	// Strings
	var stringOne rune = '1'
	var stringTwo int32 = '2'
	fmt.Println(stringOne, stringTwo)

	// Go doesn't allow automatic type promotion
	// Literals are untypd
	var numberOne int = 9
	fmt.Println(numberOne * 3) 
}