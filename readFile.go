package main

import (
	"fmt"
	"log"
	"os"
)

func ReadFile() {
	f, err := os.Open("message.txt")

	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		dta := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		fmt.Println("Read: %s\n", string(data[:n]))
	}

	//ReadFile("message.txt")
	//fmt.Println(string(content))
}
