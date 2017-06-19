package main

import (
	"flag"
	"validator"
)

func main() {
	arg := flag.String("f", "data.json", "Usage: -f=fileName.extension")
	flag.Parse()
	validator.RiseAndShine(*arg)
}
