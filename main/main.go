package main

import (
	//"validator"
	"flag"
	"fmt"
)

func main() {
	a := flag.Args
	args := ""
	for _, v := range a {
		args += v
	}

	//validator.Dance(args)
	fmt.Println(a)
}
