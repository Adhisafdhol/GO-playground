package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// You will probably not going to use this
	// This statement is considered harmful
	// This is because it could jump nearly anywhere in a program
	// But goto in go can't jump anywhere
	// Go forbids jumps that skip over variable declarations
	// Go forbids jumps that go into an inner or parallel block
	// avoid this at all cost
	// Except in rare cases

	// 	one := 1
	// 	// Cannot skip two
	// 	goto skip
	// 	two := 2
	// skip:
	// 	three := 3
	// 	fmt.Println(one, two, three)
	// 	if three > two {
	// 		// Cannot jump into an inner block
	// 		goto inner
	// 	}
	// 	if one < three {
	// 	inner:
	// 		fmt.Println("one is less than three")
	// 	}

	// Rare case
	// It is used to reduce the use of boleean flags and code duplicates
	// Try very hard to avoid goto
	attemptCount := 0
	randomNumber := rand.Intn(100)
	for randomNumber < 10000 {

		attemptCount++
		if randomNumber%3 == 0 {
			goto gotcha
		}

		if attemptCount == 10 {
			goto gameover
		}

		randomNumber = randomNumber*2 - 3

		shout := "FIND THE NUMBER THAT CAN BE DIVEDE BY 3!!!"
		fmt.Println(shout, "randomNumber: ", randomNumber)
	}

	fmt.Println("Do something when the loop completes normally")

gameover:
	fmt.Println("Sorry mate, you are DEAD!!!")

gotcha:
	if attemptCount < 10 && randomNumber%3 == 0 {
		fmt.Println("You escaped the shouting monster")
		fmt.Print("Congrats you got the number that can be divided by three")
	}
}
