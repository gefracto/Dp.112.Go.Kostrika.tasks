package validator

import (
	"fmt"
	parse "parseJsonOrXml"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"encoding/json"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
/*
func Task1Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	if (typeof(d.T1.Width) != "int") && (typeof(d.T1.Height) != "int") {
		reason = "Sides must be integers"
		ok = false
	}
	if typeof(d.T1.Symbol) != "string" {
		reason = "Symbol must be string type"
		ok = false
	}
	return
}
func Task2Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	if typeof(d.T2.Env1) != typeof(d.T2.Env2) {
		ok = false
		reason = "Envelopes must be equal type"
	}

	if (typeof(d.T2.Env1) != "task2.Envelope") || (typeof(d.T2.Env1) != "task2.Envelope") {
		ok = false
		reason = "Must be of Envelope types"
	}

	if typeof(d.T2.Env1.Side1) != "float64" || typeof(d.T2.Env1.Side2) != "float64" {
		ok = false
		reason = "Sides of envelope 1 must be float64"
	}

	if typeof(d.T2.Env2.Side1) != "float64" || typeof(d.T2.Env2.Side2) != "float64" {
		ok = false
		reason = "Sides of envelope 2 must be float64"
	}

	if d.T2.Env1.Side1 <= 0 || d.T2.Env1.Side2 <= 0 || d.T2.Env2.Side1 <= 0 || d.T2.Env2.Side2 <= 0 {
		ok = false
		reason = "Sides must be > 0"
	}
	return
}
func Task3Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task4Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task5Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task6Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task7Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
*/
func separator(n int) {
	fmt.Printf("\n\n------------->Task #%d<-------------\n\n", n)
}
func Operate(Data parse.Data){
	fmt.Println(validateData("data.json"))
	separator(1)
	fmt.Println(task1.DoTask1(Data.T1.Width, Data.T1.Height, Data.T1.Symbol))
	separator(2)
	fmt.Println(task2.DoTask2(task2.Envelope{Data.T2.Env1.Side1, Data.T2.Env1.Side2},
		task2.Envelope{Data.T2.Env2.Side1, Data.T2.Env2.Side2}))
	separator(3)
	fmt.Println(task3.DoTask3(Data.T3.SliceOfTriangles))
	separator(4)
	fmt.Println(task4.DoTask4(Data.T4.Number))
	separator(5)
	fmt.Println(task5.DoTask5(Data.T5.Min, Data.T5.Max))
	separator(6)
	fmt.Println(task6.DoTask6(Data.T6.Length, Data.T6.MaxSquare))
	separator(7)
	fmt.Println(task7.DoTask7(Data.T7.File))
}
/*
func Operate(Data parse.Data) {
	separator(1)
	if ok, reason := Task1Rules(Data); ok {
		//dowork
		fmt.Println(task1.DoTask1(Data.T1.Width, Data.T1.Height, Data.T1.Symbol))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(2)
	if ok, reason := Task2Rules(Data); ok {
		//dowork
		fmt.Println(task2.DoTask2(task2.Envelope{Data.T2.Env1.Side1, Data.T2.Env1.Side2},
			task2.Envelope{Data.T2.Env2.Side1, Data.T2.Env2.Side2}))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(3)
	if ok, reason := Task3Rules(Data); ok {
		//dowork
		fmt.Println(task3.DoTask3(Data.T3.SliceOfTriangles))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(4)
	if ok, reason := Task4Rules(Data); ok {
		//dowork
		fmt.Println(task4.DoTask4(Data.T4.Number))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(5)
	if ok, reason := Task5Rules(Data); ok {
		//dowork
		fmt.Println(task5.DoTask5(Data.T5.Min, Data.T5.Max))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(6)
	if ok, reason := Task6Rules(Data); ok {
		//dowork
		fmt.Println(task6.DoTask6(Data.T6.Length, Data.T6.MaxSquare))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	separator(7)
	if ok, reason := Task7Rules(Data); ok {
		//dowork
		fmt.Println(task7.DoTask7(Data.T7.File))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
}*/



func validateData(fileName string) (parse.DataChecking, bool, string) {
	extension := strings.Split(fileName, ".")
	var checking parse.DataChecking
	if extension[len(extension)-1] == "json" {
		contents, _ := ioutil.ReadFile("data.json")

		if b, m := parse.SimpleSpellCheckerJson(string(contents)); b == false {
			return checking, false, m
		}

		if b, m := parse.CheckForNames(string(contents)); b == false {
			return checking, false, m
		}

		json.Unmarshal(contents, &checking)

	} else if extension[len(extension)-1] == "xml" {
		contents, _ := ioutil.ReadFile("data.xml")

		if b, m := parse.SimpleSpellCheckerXml(string(contents)); b == false {
			return checking, false, m
		}

		if b, m := parse.CheckForNames(string(contents)); b == false {
			return checking, false, m
		}

		xml.Unmarshal(contents, &checking)

	} else {
		return checking, false, "Не удалось открыть файл. \nРасширение файла должно быть формата json или xml."
	}
	return checking, true, ""
}

func RiseAndShine(fileName string) (parse.Data, bool, string) {
	Data, ok, reason := parse.GetData(fileName)
	return Data, ok, reason
}
