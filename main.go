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
		item := reflect.ValueOf(&data)

		for i := 1; i <= numtasks; i++ {
			pattern := "<task%d.T%d Value>"

			if item.Elem().Field(i-1).String() != fmt.Sprintf(pattern, i, i) {
				numtasks += 1
				continue
			}

			methodname := fmt.Sprintf("Dotask%d", i)
			tools.SplitOutput(i)
			output := item.MethodByName(methodname).Call(nil)

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
