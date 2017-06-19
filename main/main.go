package main

import (
	"flag"
	"validator"
)

func main() {
	arg := flag.String("file", "data.json", "Usage: -file=fileName.extension")
	flag.Parse()
	validator.RiseAndShine(*arg)
}
