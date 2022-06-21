package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	bt "internal/baiton"
)

var opts struct {
	Delay   uint   `short:"d" value-name:"D" long:"delay" default:"1" description:"Wait for D seconds between checks"`
	Timeout uint   `short:"t" value-name:"T" long:"timeout" default:"60" description:"Fail if file did not appear in T seconds"`
	File    string `short:"f" value-name:"F" long:"file" default:"" description:"Wait for this file" required:"true"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	file := opts.File
	d := opts.Delay
	t := opts.Timeout

	result := bt.WaitFile(file, d, t)
	if result != nil {
		fmt.Printf("ERROR: %v\n", result)
		os.Exit(2)
	}

	fmt.Printf("file %v is ready\n", file)
}
