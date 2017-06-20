package main

import (
	"flag"
	"fmt"
	pars "parseJsonOrXml"
)

func main() {

	arg := flag.String("file", "data.json", "Usage: -file=fileName.extension")
	flag.Parse()
	data, ok, reason := pars.MakeAStructFromJson(*arg)
	if ok {
		pars.Operate(data)
	} else {
		fmt.Println(reason)
	}


}
