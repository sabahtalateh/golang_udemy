package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file name is not specified.")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		bs := make([]byte, 1024)
		for {
			n, err := file.Read(bs)
			if n == 0 {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(bs))
		}
	}
}
