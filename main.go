package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("n", 1050, "number")
	flag.Parse()

	fmt.Printf("Hello, %v\n", *num)
}
