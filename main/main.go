package main

import (
	"flag"
	"fmt"
	"validator"
)

func main() {
	arg := flag.String("file", "data.json", "Usage: -file=fileName.extension")
	flag.Parse()
	data, ok, reason := validator.RiseAndShine(*arg)
	if ok {
		validator.Operate(data)
	} else {
		fmt.Println(reason)
	}
}
