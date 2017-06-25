package main

import (
	"flag"
	"fmt"
	"reflect"
	"tools"
)



func main() {
	arg := flag.String("file", "data.json", "Usage: -file=fileName.extension")
	flag.Parse()

	if data, err := tools.GetData(*arg); err == nil {
		var QuantityOfTasks int = reflect.ValueOf(&data).Elem().NumField()
		for i := 1; i <= QuantityOfTasks; i++ {
			methodname := fmt.Sprintf("Dotask%d", i)
			tools.SplitOutput(i)
			reflect.ValueOf(&data).MethodByName(methodname).Call(nil)
		}

	} else {
		fmt.Println(err)
	}
}