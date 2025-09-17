package main

import (
	"errors"
	"fmt"
)

func div(number int, divisor int) (result int, err error) {
	if divisor == 0 {
		return result, errors.New("cannot divide by zero")
	}

	return number / divisor, nil
}

func createPrefixer(prefix string) func(word string) string {
	return func(word string) string {
		return prefix + word
	}
}

func main() {
	resultOne, _ := div(65, 6)
	fmt.Println(resultOne)
	resultTwo, err := div(57, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultTwo)
	}

	prefixFn := createPrefixer("HELLO!")
	fmt.Println(prefixFn("Jack"))
}
