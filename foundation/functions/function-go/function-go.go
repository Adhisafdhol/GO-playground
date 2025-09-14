package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
// Function declaration
func name(param type) return type {
	body
}

*/

/*
	The consecutive input parameters with the same type can be written

as: func divide(num, divisor int)
but i personally think it is less intuitive
*/
func divide(num int, divisor int) int {
	if divisor == 0 {
		return 0
	}

	return num / divisor
}

// Go doesn't have named and optional parameters
// You can simulate optional parameters with a atruct
type Person struct {
	name  string
	age   int
	quote string
}

func logPersonQuote(person Person) {
	if person.quote != "" {
		fmt.Println(person.quote)
	} else {
		fmt.Println(person.name, "doesn't have a personal quote")
	}

}

// Variadic input parameters in Go
// The variadic parameter must be the last or only parameter
// The crated variable is a slice

func addTo(base int, toAdd ...int) int {
	for _, v := range toAdd {
		base += v
	}

	return base
}

func consoleMany(values ...string) {
	for _, v := range values {
		fmt.Println(v)
	}
}

// Go allow multiple return values
/* Unlike Python, Go multiple return values are not destructured
tupple's values. You must assign it to multiple variables */
func divideAndRemainder(num int, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / divisor, num % divisor, nil
}

// Go allow named return values
/* What you are doing is predeclaring variables that you use within
function the function to hold the return values.
They are initialized to their zero values and need to be inside
parentheses. you can use _ to name some values only.
*/
func addAndSubstract(a int, b int) (sum int, difference int) {
	sum, difference = a+b, a-b

	return sum, difference
}

// Go also allow blank return
// But do not do this
/* The returned variables are the last values set to the named return
variables or the initial zero values */
func addAndSubstractBlank(a int, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return
}

func main() {
	result := divide(9, 4)
	fmt.Println(result)

	theo := Person{
		name: "Theo",
		age:  32,
	}
	logPersonQuote(theo)

	base := 10
	// Should we check side effects?
	resultTwo := addTo(base, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultTwo)
	// It seems like the original base is untouched
	// interesting
	fmt.Println(base)
	numbers := []int{2, 4, 6, 8, 10, 12}
	resultThree := addTo(base, numbers...)
	fmt.Println(resultThree)
	// YOu can even supply zero value for the last parameter
	resultFour := addTo(base)
	consoleMany()
	fmt.Println(resultFour)
	_, _, errorFive := divideAndRemainder(1, 0)
	if errorFive != nil {
		fmt.Println(errorFive)
	}

	sumResult, differenceResult := addAndSubstract(1, 2)
	fmt.Println(sumResult, differenceResult)
	a, b := addAndSubstractBlank(3, 6)
	fmt.Println(a, b)

	// Functions are values
	/* The type of a function is built out of the keyword func,
	its parameter types,and return value types. This is its signature */
	// func type default to nil
	var sayHi func(string) string
	fmt.Println(sayHi == nil)
	sayHi = func(name string) string {
		fmt.Println("HI", name)
		return "DONE"
	}
	sayHi("JACK")
	stringOne := "1"
	one, _ := strconv.Atoi(stringOne)
	fmt.Println(one + 1)
	// Functions type declaration
	type multiply func(int, int) int
	// you cannot name a function declared with this syntax
	screamHello := func() {
		fmt.Println("HELLO!")
	}
	screamHello()
	// You can also immediately invoke it without assigning it to a var
	func() {
		fmt.Println("Blimey!")
	}()
	createMultiplyFunc := func() multiply {
		multiply := func(a int, b int) int {
			return a * b
		}

		return multiply
	}

	multiplyFunc := createMultiplyFunc()
	resultMultiply := multiplyFunc(3, 5)
	fmt.Println(resultMultiply)
}
