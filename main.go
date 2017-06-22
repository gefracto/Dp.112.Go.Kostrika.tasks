package main

import (
	"flag"
	"fmt"
	"parser"
)

func main() {

	arg := flag.String("file", "data.json", "Usage: -file=fileName.extension")
	flag.Parse()
	data, ok, reason := parser.GetData(*arg)
	if ok {
		parser.Operate(data)
	} else {
		fmt.Println(reason)
	}

}
