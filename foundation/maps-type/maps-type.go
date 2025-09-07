package main

import (
	"fmt"
	"maps"
)

func main() {
	// Declare a map with string keys and int values
	// key:value
	// the zero value is nil
	var nilMap map[string]int
	// You cannot write to a nil map and it will cause a panic
	fmt.Println("nilMap:", nilMap)
	fmt.Println(len(nilMap))

	// Create a map literal
	// You can read and write using this syntax
	literalMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("literalMap:", literalMap)

	// Create a map with a default size
	/* Make still creates a map with 0 length and it can grow past the specified
	size */
	defaultSizeMap := make(map[int]string, 3)
	fmt.Println("defaultSizeMap:", defaultSizeMap)
	fmt.Println(len(defaultSizeMap))
	// read and write a map
	defaultSizeMap[1] = "one"
	fmt.Println("defaultSizeMap:", defaultSizeMap)
	fmt.Println(defaultSizeMap[1])

	// Use the comma ok idiom to check if the key exists
	/* This differentiates whether it reads the key's value or the zero
	value of non-existent key */
	notExist, ok := defaultSizeMap[3]
	fmt.Println("notExist, ok:", notExist, ok)

	exist, ok := defaultSizeMap[1]
	fmt.Println("exist, ok:", exist, ok)

	// Delete from maps: delete(map, key)
	// nothing happens if the key doesn't exist or when the map is nil
	// Doesn't return anything
	delete(defaultSizeMap, 1)
	fmt.Println("defaultSizeMap:", defaultSizeMap)

	// Emptying a map
	clear(defaultSizeMap)
	fmt.Println("defaultSizeMap:", defaultSizeMap, len(defaultSizeMap))
	// Compare maps with maps.Equal or maps.EqualFunc
	mapToCompareOne := map[int]string{
		1: "one"}
	mapToCompareTwo := map[int]string{
		2: "two"}
	fmt.Println(maps.Equal(mapToCompareOne, mapToCompareTwo))

	// Simulate a set with a map
	set := map[int]bool{}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}

	for _, value := range values {
		// The same value will result in setting the key into true once more
		set[value] = true
	}

	fmt.Println("set:", set)
	fmt.Print("values let", len(values), "set len", len(set))
	fmt.Println(set[3])
	fmt.Println(set[9])
	if set[9] {
		fmt.Println("9 is in the set")
	}
}
