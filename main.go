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

	err, data := tools.GetData(*arg)
	if err != nil {
		fmt.Println(err)
	} else {

		var numtasks int = reflect.ValueOf(&data).Elem().NumField()

		for i := 1; i <= numtasks; i++ {
			methodname := fmt.Sprintf("Dotask%d", i)
			tools.SplitOutput(i)
			output := reflect.ValueOf(&data).MethodByName(methodname).Call(nil)

			Err := output[0]
			Dat := output[1]

			if reflect.Value.IsNil(Err) {
				fmt.Println(Dat)

			} else {
				fmt.Println(Err)
			}
		}

	}
}
