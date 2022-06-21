package main

import (
	"flag"
	"fmt"
	bt "internal/baiton"
)

func main() {
	num := flag.Int("n", 1050, "number")
	flag.Parse()
	r := bt.Add(*num, *num)

	fmt.Printf("Hello, %v\n", r)
}
