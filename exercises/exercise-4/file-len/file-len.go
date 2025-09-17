package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data := make([]byte, 1)
	var strings = []byte{}

	for {
		count, err := file.Read(data)

		if err == nil {
			strings = append(strings, data[:count][0])
		}

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			// end of file
			break
		}
	}

	fmt.Println(len(strings), " bytes long")

}
