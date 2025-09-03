package main

import "fmt"

func main() {

	var numberOne int = 10
	fmt.Println(
		"Declare a variable with explicit type\n",
		"numberOne:",
		numberOne,
	)
	var numberTwo = 10
	fmt.Println(
		"Declare variable with implicit type if the left side of the '='"+
			" matches the default type\n",
		"numberTwo:",
		numberTwo,
	)
	var numberThree int
	fmt.Println(
		"Declare a variable and assign it the zero value\n",
		"numberThree:",
		numberThree,
	)
	var numberFour, numberFive int
	fmt.Println(
		"Declare multiple variables once with var\n",
		"numberFour & numberFive:",
		numberFour,
		numberFive,
	)
	var numberSix, stringOne = 15, "Fifteen"
	fmt.Println(
		"Declare multiple variables of different variables once with var\n",
		"numberSix & stringOne:",
		numberSix,
		stringOne,
	)
	var (
		numberSeven          int
		numberEight          = 20
		numberNine           = 25
		numberTen, stringTwo = 30, "GREAT DAY!"
	)
	fmt.Println(
		"Declare multiple variables with declaration list\n",
		"numberSix & stringOne:",
		numberSeven,
		numberEight,
		numberNine,
		numberTen,
		stringTwo,
	)

	fmt.Println("numberTen:", numberTen)
	numberTen = numberTen * 2
	fmt.Println("numberTen:", numberTen)
	// Short declaration and assignment format with :=
	// It allows you to assign values to existing variables but only
	// when at least one new variable is on the left side of the :=
	// := is not legal outside of function
	shortNumOne := 1
	shortNumOne, shortNumTwo := 2, 3
	fmt.Println(
		"Declare variables wtih short declaration :=\n",
		shortNumOne,
		shortNumTwo,
	)

	// Constanst in Go are a way to give names to literals.
	// No way to declare that a variable is immutable
	const constantNumOne uint32 = 10
	const constantStringOne = "hello"
	fmt.Println(
		"Declare a variable with constant\n",
		constantNumOne,
		constantStringOne,
	)
}
