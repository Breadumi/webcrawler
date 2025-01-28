package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 1 {
		log.Fatal("no website provided")
	}
	if len(args) > 1 {
		log.Fatal("too many arguments provided")
	}
	BASE_URL := args[0]
	fmt.Printf("starting crawl of: %v\n", BASE_URL)
}
