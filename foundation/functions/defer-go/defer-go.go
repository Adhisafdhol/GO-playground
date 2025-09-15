package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Cleanup code is attached to the function with the defer keyword
	// It is used to cleanup temporary resources
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal((err))
	}
	// Ensure the cleanup code runs
	// Defer delay function invocation until the surrounding function exists
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	// You can defer multiple functions in a Go function
	// They run in LIFO order
	// The code within defer function runs after the return statement
	/* You can supply a function that returns values to defer But there's
	no way to read those values. */
	/* A deferred functioon can examine or modify the return values of
	its surrounding function*/
	testDefer := func(a int, b int) {
		defer func() {
			if err == nil {
				fmt.Println("SUCCESS")
			}

			if err != nil {
				fmt.Println("FAIL")
			}
		}()

		_, err = func(a int, b int) (difference int, err error) {
			if a == 0 {
				err = errors.New("a is zero")
			}

			if b == 0 {
				err = errors.New("b is zero")
			}

			if b != 0 {
				difference = a / b
			}

			return difference, err
		}(a, b)
	}

	testDefer(20, 20)

	/* A common pattern in Go for a function that allocates a resource to
	   return a closure that cleans up the resource */
	// for example a function that returns data, clenup, error
}
