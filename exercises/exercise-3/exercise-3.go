package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := make([]int, 0, 100)
	fmt.Println(len(numbers), cap(numbers))

	for len(numbers) < 100 {
		randomNumber := rand.Intn(100)
		numbers = append(numbers, randomNumber)
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	for _, v := range numbers {
		fmt.Println("number: ", v)
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println("inner total(shadowing): ", total)
	}

	fmt.Println("outer total: ", total)

}
