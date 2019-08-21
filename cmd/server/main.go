package main

import (
	"fmt"
	"log"
)

var commit string

func main() {
	log.SetFlags(0) // this removes timestamp prefixes from logs
	fmt.Printf("Hello world: %v\n", commit)
}
