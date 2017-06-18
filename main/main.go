package main

import (
	"validator"
	"os"
)

func main() {
	args := string(os.Args)
	validator.Dance(args)
}
