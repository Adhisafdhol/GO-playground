package main

import "fmt"

// Go only has 25 keywords.
/* The built-in types, constants, and functions aren't included in that
list. */
// These are predeclared identifiers defined in the universe block.
// The universe block contains all other blocks.
// Which means they can be shadowed!!
/* Be very carefull ... or you would accidentally redefined an
identifier in the universe block ...*/

func main() {
	// Don't do this!
	bool := "ERROR"
	fmt.Println(bool)

	// How does block work in go?
	// var error bool = "ERROR"
	// fmt.Print(error)
	number := 27

	if number == 27 {
		fmt.Println(number)
		// Shadowing variables
		/* declare a variable with the same name as a variable in the
		outer block */
		// Go vet won't report shadowing as an error
		number := 17
		/* This will print 17 but doesn't change the number ourside of this
		block */
		// Be careful to not shadow a package import
		// fmt := "hi";
		fmt.Println(number)
		number -= 27
		fmt.Println(number)
	}

	fmt.Println(number)

	var numberTwo int = 5

	if numberTwo >= 5 {
		// var numberTwo = 9
		numberTwo, numberThree := 1, 2
		fmt.Println(numberTwo, numberThree)
	}

	fmt.Println(numberTwo)
}
