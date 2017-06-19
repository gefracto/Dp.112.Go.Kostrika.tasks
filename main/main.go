package main

import (
	"flag"
	"validator"
)

func main() {
	arg := flag.String("file", "data.xml", "Usage: -file=fileName.extension")
	flag.Parse()
	validator.RiseAndShine(*arg)
}
