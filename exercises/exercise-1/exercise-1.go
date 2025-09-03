package main

import "fmt"

func main() {
	var intOne int = 3
	var floatOne float64 = float64(intOne)
	fmt.Println(intOne, floatOne)

	// typed constant can only be assigned to a variable of the same type
	// Use untyped constant
	const value = 10
	var intTwo int = value
	var floatTwo float64 = value
	fmt.Println(intTwo, floatTwo)

	var maxByteOne byte = 255
	var maxInt32One int32 = 2147483647
	var maxUint64 uint64 = 18446744073709551615
	fmt.Println(maxByteOne, maxInt32One, maxUint64)
	maxByteOne += 1
	maxInt32One += 1
	maxUint64 += 1
	fmt.Println(maxByteOne, maxInt32One, maxUint64)
}