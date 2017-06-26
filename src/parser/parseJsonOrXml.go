package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
)

type Data struct {
	T1 struct {
		Width, Height int
		Symbol        string
	}
	T2 struct {
		Env1, Env2 task2.Envelope
	}
	T3 struct {
		SliceOfTriangles []task3.Triangle
	}
	T4 struct {
		Number int
	}
	T5 struct {
		Min, Max int
	}
	T6 struct {
		Length, MaxSquare int
	}
	T7 struct {
		File string
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

func separator(n int) {
	fmt.Printf("\n\n------------->Task #%d<-------------\n\n", n)
}


func Operate(Data Data) {
	separator(1)
	if ok, data, reason := task1.DoTask1(Data.T1.Width, Data.T1.Height, Data.T1.Symbol); ok {
		fmt.Println(data)
	} else {fmt.Println(reason)}

	separator(2)
	if ok, data, reason := task2.DoTask2(task2.Envelope{Data.T2.Env1.Side1, Data.T2.Env1.Side2},
		task2.Envelope{Data.T2.Env2.Side1, Data.T2.Env2.Side2}); ok {
		fmt.Println(data)
	} else {fmt.Println(reason)}

	separator(3)
	if ok, data, reason := task3.DoTask3(Data.T3.SliceOfTriangles); ok {
		fmt.Println(data)
	} else {fmt.Println(reason)}

	separator(4)
	if ok, data, reason := task4.DoTask4(Data.T4.Number); ok {
		fmt.Println(data, ok)
	} else {fmt.Println(ok, reason)}

	separator(5)
	if ok, data, reason := task5.DoTask5(Data.T5.Min, Data.T5.Max); ok {
		fmt.Println(data)
	} else {fmt.Println(reason)}

	separator(6)
	if ok, data, reason := task6.DoTask6(Data.T6.Length, Data.T6.MaxSquare); ok {
		fmt.Println(data)
	} else {fmt.Println(reason)}

	separator(7)
	if ok, data, reason := task7.DoTask7(Data.T7.File); ok {
		fmt.Println(data)
	} else {fmt.Println(reason, ok)}
}
