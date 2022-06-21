package main

import (
	"flag"
	"fmt"
	"os"

	bt "internal/baiton"
)

var filepath = flag.String("file", "", "Wait for this file")
var delay = flag.Int("delay", 1, "Wait for D seconds between checks")
var timeout = flag.Int("timeout", 60, "Fail if file did not appear in T seconds")

func checkArgs() bool {
	if *filepath == "" {
		fmt.Println("ERROR: file not specified")
		return false
	}
	if *delay < 0 {
		fmt.Println("ERROR: delay should be >0")
		return false
	}
	if *timeout < 0 {
		fmt.Println("ERROR: timeout should be >0")
		return false
	}
	return true
}

func main() {
	flag.Parse()
	checkArgs()

	if !checkArgs() {
		os.Exit(1)
	}

	file := *filepath
	d := *delay
	t := *timeout

	result := bt.WaitFile(file, d, t)
	if result != nil {
		fmt.Printf("ERROR: %v\n", result)
		os.Exit(2)
	}

	fmt.Printf("file %v is ready\n", file)
}
