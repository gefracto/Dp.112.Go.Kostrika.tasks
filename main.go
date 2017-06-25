package main

import (
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"flag"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	"strings"
	"strconv"
	"reflect"
)

type Data struct {
	task1.T1
	task2.T2
	task3.T3
	task4.T4
	task5.T5
	task6.T6
	task7.T7
}

func main() {
	arg := flag.String("file", "data.xml", "Usage: -file=fileName.extension")
	flag.Parse()
	data, _,_ := GetData(*arg)

	var Quantityoftasks int = 7
	for i:=1; i<=Quantityoftasks; i++ {
		nameofmethod := "Dotask" + strconv.Itoa(i)
		reflect.ValueOf(&data).MethodByName(nameofmethod).Call(nil)
	}

}



func GetData(fileName string) (Data, bool, string) {
	var MyData Data
	if strings.HasSuffix(fileName, ".json") {
		var contents []byte
		var err error
		if contents, err = ioutil.ReadFile(fileName); err != nil {
			return MyData, false, "Не удалось найти файл"
		}
		if err = json.Unmarshal(contents, &MyData); err != nil {
			return MyData, false, "Не удалось распаковать json"
		}

	} else if strings.HasSuffix(fileName, ".xml") {

		var contents []byte
		var err error
		if contents, err = ioutil.ReadFile(fileName); err != nil {
			return MyData, false, "Не удалось найти файл"
		}
		if err = xml.Unmarshal(contents, &MyData); err != nil {
			return MyData, false, "Не удалось распаковать xml"
		}

	} else {
		return MyData, false, "Не удалось открыть файл. " +
			"\nРасширение файла должно быть формата json или xml."
	}

	return MyData, true, ""
}