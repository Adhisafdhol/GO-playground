package main

import "fmt"

func main() {
	greetings := []string{"hello", "hola", "नमस्ते", "こんにちわ", "Привет"}
	subsliceGreetingsOne := greetings[:2]
	subscliceGreetingsTwo := greetings[1:4]
	subsliceGreetingsThree := greetings[3:]

	fmt.Println(subsliceGreetingsOne)
	fmt.Println(subscliceGreetingsTwo)
	fmt.Println(subsliceGreetingsThree)

	message := "hi 👩🏻 and 👦🏻"
	messageRunes := []rune(message)

	fmt.Println(string(messageRunes[3]))

	type employee struct {
		firstName string
		lastName string
		id int
	}

	anita := employee{
		id: 123,
	}

	bonnie := employee{
		firstName: "Bonnie",
		lastName: "Johnson",
		id: 133,
	}

	var clark = employee{
	}

	clark.firstName = "Clark"
	clark.lastName = "Donnavan"
	clark.id = 233

	fmt.Println(anita, bonnie, clark)

}
