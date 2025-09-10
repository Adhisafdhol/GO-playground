package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := 9

	/* Any variable declared within the braces of an if or else statements
	exists only within that block */
	if number == 0 {
		fmt.Println("Number is 0")
	} else if number > 6 {
		fmt.Println("Number's bigger than 6")
	} else {
		fmt.Println("Number's between 0 and 6")
	}

	/* You can declare variables that are scoped to the condition and
	if and else block */
	// Shadowing is still applied to these variables
	if specialNumber := rand.Intn(10); specialNumber == 0 {
		specialNumber := "YOLO"
		fmt.Println(specialNumber)
		fmt.Println("That's too low")
	} else if specialNumber > 6 {
		fmt.Println("Perfection!")
	}

	// The special variable won't exist outside of the the previous
	// if and else's block and condition
	// fmt.Println(specialNumber)
}
