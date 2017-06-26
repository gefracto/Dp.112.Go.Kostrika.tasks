package main

import (
	"flag"
	"fmt"
	"reflect"
	"tools"
)

func main() {
	arg := flag.String("file", "data.json",
		"Usage: -file=fileName.extension")

	flag.Parse()

	if data, err := tools.GetData(*arg); err == nil {
		var QuantityOfTasks int = 3 //reflect.ValueOf(&data).Elem().NumField()
		for i := 1; i <= QuantityOfTasks; i++ {
			methodname := fmt.Sprintf("Dotask%d", i)
			tools.SplitOutput(i)
			output := reflect.ValueOf(&data).MethodByName(methodname).Call(nil)

			if reflect.Value.IsNil(output[0]) {
				fmt.Println(output[1])
			} else {fmt.Println(output[0])}
		}

	} else {
		fmt.Println(err)
	}
}