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
	// The update also effects the initial slice
	sliceSliceOne[0] = 999
	fmt.Println(sliceOne)
	// Confusing slice behaviours with append
	sliceAOne := []string{"a", "b", "c", "d"}
	subSliceAOne := sliceAOne[:2]
	fmt.Println("sliceAOne:", sliceAOne)
	fmt.Println("subSliceAOne", subSliceAOne)

	fmt.Println(cap(sliceAOne), cap(subSliceAOne))
	/* Append doesn't work as expected due to sublice being created with
	the capacity minus the offset and share the same memory*/
	/* Which means they same the same memory up to the capacity of rhe original
	slice, after the capacity has ben reached it appends normally */
	subSliceAOne = append(subSliceAOne, "z")
	subSliceAOne = append(subSliceAOne, "za")
	// sliceAOne[3] = "STRIKE"
	// sliceAOne[0] = "STRIKE"
	// subSliceAOne[0] = "ACROBAT"
	fmt.Println("sliceAOne:", sliceAOne)
	fmt.Println("subSliceAOne", subSliceAOne)
	subSliceAOne = append(subSliceAOne, "zb")
	subSliceAOne = append(subSliceAOne, "zc")
	fmt.Println("sliceAOne:", sliceAOne)
	fmt.Println("subSliceAOne", subSliceAOne)
	/* It seems like the connection is broken after appending outside the original
	slice capacity */
	sliceAOne[3] = "STRIKE"
	sliceAOne[0] = "STRIKE"
	fmt.Println("sliceAOne:", sliceAOne)
	fmt.Println("subSliceAOne", subSliceAOne)
	subSliceAOne[2] = "ACROBAT"
	fmt.Println("sliceAOne:", sliceAOne)
	fmt.Println("subSliceAOne", subSliceAOne)
	// Use the full slice expression to prevent this to limit its capacity
	var subSliceATwo = sliceAOne[0:2:2]
	fmt.Println(subSliceATwo)
	subSliceATwo = append(subSliceATwo, "NOPE")
	fmt.Println(subSliceATwo)
	// This prevent the original slice from sharing the whole capacity memory
	fmt.Println(sliceAOne)
	// Copy slice with copy(destination[start:end], source[start:end])
	var destination = make([]int, 2)
	var source = []int{1, 2, 3, 4}
	copy(destination, source)
	fmt.Println("copy:", destination)
	// copy overlapping section
	// YOu can use copy with array
	copyOne := copy(source[:3], source[1:])
	fmt.Println(copyOne, source)
	// You can slice arrays
	// To conver an entire array into a slice use the [:] syntax
	// [start:end]
	/* You can convert a subset of an array into a slice which has the same 
	memory problems as a subslice */

	// Convert slicest to arrays
	// This creates a copy of the slice
	// Arr type convertion doesn't work with [...] syntax
	/* The size of the array can be smaller than the size of the slice but not
	bigger and will cause panic at runtime */
	arrFromSlice := [4]int(sliceOne)
	fmt.Println((arrFromSlice))
}
