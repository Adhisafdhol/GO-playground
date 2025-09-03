package main

import (
	"fmt"
	"slices"
)

func main() {
	// Array
	var arrayOne [3]int
	var arrayTwo = [3]int{10, 20, 30}
	// Sparse array
	var arrayThree = [12]int{1, 3: 3, 6, 9: 9, 81}
	var arrayFour = [...]int{3, 6, 9}
	fmt.Println(arrayOne, arrayTwo, arrayThree, arrayFour)
	// Compare arrays with == or !=
	var arrayOneCopy = arrayOne
	fmt.Println(arrayOne == arrayFour)
	fmt.Println(arrayOne == arrayOneCopy)
	// Simulate two dimensional array
	var array2DOne [3][3]int
	fmt.Println((array2DOne))
	// Access array's item
	array2DOne[0][0] = 1
	fmt.Println(array2DOne[0][0])
	fmt.Println(array2DOne)
	/* Cannot read or write past the end of an array or use a negative
	index */
	// Go consider the size of the array to be part of the type
	/* Can't use type conversion to directly convert arrays of differet
	sizes to identical types */
	// Check array's length
	fmt.Println(len(arrayOne))
	// Slice
	var sliceOne = []int{10, 20, 30}
	var sliceTwo = []int{1, 3: 3}
	var slice2DOne [][]int
	var sliceThree []int
	fmt.Println("Length", len(sliceTwo))
	fmt.Println(sliceOne, sliceTwo, slice2DOne, sliceThree)
	fmt.Println(sliceThree == nil)
	fmt.Print(slice2DOne == nil, "\n")
	// Slice type isn't comparable using == and !=
	fmt.Println(slices.Equal(sliceOne, sliceTwo))
	// Grow slices with append
	sliceTwo = append(sliceTwo, 3)
	fmt.Println(sliceTwo)
	sliceTwo = append(sliceTwo, 4, 5, 6)
	fmt.Println(sliceTwo)
	// Append size onto another slice
	sliceOne = append(sliceOne, sliceTwo...)
	fmt.Println(sliceOne)
	/* Each slice has a capacity which is the number of consecutive memory
	reserved */
	/* Go runtime increases slice by more than one each time it runs out
	of capacity using a special rule*/
	fmt.Println(len(sliceOne), cap(sliceOne))
	sliceOne = append(sliceOne, 1)
	fmt.Println(len(sliceOne), cap(sliceOne))
	sliceOne = append(sliceOne, 1)
	fmt.Println(len(sliceOne), cap(sliceOne))
	/* You can create slice with specified length, type, and capacity with
	the make function */
	var sliceFour = make([]int, 5, 10)
	fmt.Println(sliceFour, len(sliceFour), cap(sliceFour))
	sliceFive := []string{"one", "two", "three"}
	fmt.Println(sliceFive)
	// Clear funtion sets all of the slice's elements to their zero value.
	// The lenght of the  slice remains unchanged
	clear(sliceFive)
	fmt.Println(sliceFive, cap(sliceFive), len(sliceFive))
	// slicing slices
	// a Slice expressios creates a slice from a slice
	// [starting offset:ending offset]
	// if the ending offset is index 5 it will include up to index 4
	var sliceSliceOne = sliceOne[3:6]
	fmt.Println(sliceSliceOne)
	// A slice of slice shares the memory with the original slice
	// Which means the changes in the original slice will affect it.
	sliceOne[3] = 1
	sliceOne[4] = 2
	sliceOne[5] = 3
	fmt.Println(sliceSliceOne)
}
