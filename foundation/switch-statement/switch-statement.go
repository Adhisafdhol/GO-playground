package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range numbers {
		// switch scopedVariable;toCompare
		switch number := v; number {
		// Cases in switch statements in Go don't fall through
		// No need for break, YAY!
		// Empty cases means nothing happens
		/* Go allows switch to continue with fallthrough keyword, think before
		using it */
		// YOu can still use break keyword to break out of the case
		case 1, 2:
			fmt.Println("Value is less than 3")
		case 3, 4:
			fmt.Println("Value is less than 5")
		case 5, 6, 7, 9:
			fmt.Println("Value is less than 10")
		default:
			fmt.Println("Values is unknown")
		}
	}

	// Blank switch
	// You can leave out parts from a switch statement
	// leave out toCompare, this allows you to use any boolean comparison
	for _, number := range numbers {
		switch {
		case number*3 < 10:
			fmt.Println("Number x 3 is less than 10")
		case number*3 < 20:
			fmt.Println("Number x 3 is less than 20")
		default:
			fmt.Println("Number * 3 i bigger than 20")
		}
	}
}
