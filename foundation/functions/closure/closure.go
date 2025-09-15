package main

import (
	"fmt"
	"sort"
)

func main() {
	// What are closures?
	/* A CS term that means that functions declared inside functions are
	able to access and modify variables declared in the outer function */
	// The inner function captures its outside functions'r variables
	number := 9
	fn := func() {
		fmt.Println(number)
		number = number * 3
	}
	fn()
	fmt.Println(number)

	// A sample use case for closures
	type timesByNum func(int) int
	createTimesByNum := func(toMultiply int) timesByNum {
		return func(a int) int {
			return a * toMultiply
		}
	}
	// This function remembers its outer function variable and argument
	timesBy5 := createTimesByNum(5)
	fmt.Println(timesBy5(3))
	// Passing functions as parameters
	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	people := []Person{
		{"Andy", "James", 31},
		{"Brandon", "Thomas", 32},
		{"Connor", "Anderson", 36},
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	fmt.Println(people)
	/* Higher-order functions are functions that have functions as inputs
	or return values */
}
