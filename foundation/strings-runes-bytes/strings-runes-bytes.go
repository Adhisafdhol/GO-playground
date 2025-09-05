package main

import "fmt"

func main() {
	// Go uses a sequence of bytes to represent a string
	/* The bytes don't have to be in any particular character encoding, but UTF-8
	is the standard
	*/

	var hello string = "HELLO 🎃"
	var h byte = hello[0]
	fmt.Println("hello:", hello)
	fmt.Println("h", h)
	/* Slice expression notation that is used with arrays and slices also works
	with strings*/
	// Strings are immutable
	var hel string = hello[:3]
	fmt.Println("hel:", hel) 
	// Doesn't work with emoji or any code point that is bigger than one byte
	var emoji = hello[5:7]
	fmt.Println("emoji:", emoji)
	// Check the string length in bytes
	fmt.Println(len(hello))
	
	// Type conversions
	/* Go blocks a type conversion to string from any integer type other than rune
	or byte since Go 1.15 */
	// Type conversion from int 65 results in the character "A" instead of "65"
	var runeOne = '1'
	var stringOne = string(runeOne)
	var byteOne byte = '1'
	var stringTwo string = string(byteOne)
	fmt.Println(stringOne, stringTwo)
	/* String can be converted back and forth to a slice of bytes or a slice of 
	runes */
	var byteSlice []byte = []byte(hello)
	var runeSlice []rune = []rune(hello)
	fmt.Println("byteSlice:", byteSlice)
	fmt.Println("runeSlice:", runeSlice)
	// Most data in Go is read and written as a sequence of bytes
	/* The most common string type conversions are back and forth with a slice
	of bytes */
	// UTF-8 <The most common encoding for unicode>
	// The two Creatorors of Go, Ken Thomson and Rob Pike invented UTF-8 in 1992
	// Very interesting

}