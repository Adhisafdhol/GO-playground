package main

import (
	"fmt"
)

func main() {
	// Loop in go

	// The complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for initialisation;condition;increment
	// Use := in the initialisation, the var is not allowed here.
	// Shadowing is applied in the for loop
	/* You can leave out one or more of the three parts of the for
	statement */

	i := 0

	for ; i < 100; i++ {
		fmt.Println(i * 3 / 2)
	}

	// The condition-only for statement
	// Do not include semicolons on this kind of statement
	// This acts like a while function
	indexTwo := 0

	for indexTwo < 100 {
		if indexTwo%2 == 0 {
			indexTwo++
			fmt.Println("EVEN!!!!")
		} else {
			if indexTwo%3 == 0 {
				indexTwo += 2
			} else {
				indexTwo++
			}

			fmt.Println("ODD!!!!")
		}
	}

	// The infinite for statement

	// This is an infinite loop that loops forever
	// for {
	// 	fmt.Println("HI THERE T!!")
	// }

	// break and continue

	// break statement exists the loop immediately

	for {
		i := 0

		fmt.Println("We interrupt this program to show you!")
		if i == 0 {
			break
		}
	}

	/* continue keyword skips over the rest of the for loop's body and
	proceeds directly to the next iteration */

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Can be divided with 2")
			continue
		}

		if i%3 == 0 {
			fmt.Print("Can be divided with 3\n")
			continue
		}

		fmt.Print("Can be divided with 1\n")
	}

	// For-range statement
	// is used for iterating over elements in some of Go's built in types
	/* You can only use this to iterate over the built-in compound types
	and user-defined types that are based on them */

	numbers := []int{1, 2, 3, 4, 5, 6}

	/* use i which is the short form of index when iterating over
	a slice, array, or string */
	// use k which is the short form of key when iterating over a map
	for i, v := range numbers {
		fmt.Println("index: ", i)
		fmt.Println("value: ", v)
	}

	// Use _ for unused variables in the for range loop
	for _, v := range numbers {
		fmt.Println("Value * 3: ", v*3)
	}

	// You can leave off the second variable if you don't use the value
	for i := range numbers {
		fmt.Println("NINE TIMES THE INDEX??? ", i*9)
	}

	numberSet := map[int]bool{1: true, 2: true, 3: true}

	for k := range numberSet {
		fmt.Println("key: ", k)
	}

	// Use underscore to hide value that you want to ignore

	// Interesting for-range iteration over a map

	names := map[string]int{
		"Anton":    31,
		"Berlin":   32,
		"Clarkson": 33,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Index", i)

		for k, v := range names {
			fmt.Println(k, v)
		}
	}

	// The output of the above loop varies
	// This is a security feature
	// This prevent people to write code that assumed the order was fixed
	// Prevent Hash DoS attack
	/* With exception of the formatting functions (like fmt.Printlng)
	That always output pamp with their keys in ascending sorted order */

	// Iteraing over strings
	stringOne := "HELLO 🧛🏻‍♂️"

	// For loop iterates over the runes not the bytes
	/* When it encounters a miltybite rune it converts UTF-8
	representation into a single 32-bit number and assigns it to the
	value. */
	// Invalid byte is replaced with hex value 0xfffd
	for i, rune := range stringOne {
		fmt.Println("rune:", rune)
		fmt.Println(i, string(rune))
	}

	numbersSlice := []int{1, 2, 3}

	// for range value is a copy
	for _, v := range numbersSlice {
		// This doesn't affect the original slice
		v *= 2
	}

	fmt.Println(numbersSlice)

	// Labeling for statements

	stringsSlice := []string{"HELLO", "IT's", "ME", "MARIO!"}
outer:
	for _, v := range stringsSlice {
		for i, r := range v {
			fmt.Println(i, string(r))
			if r == '\'' {
				/* Use the label to apply break or continue on that label
				scope */
				break outer
			}
		}
	}

}
