package validator

import (
	"fmt"
	parse "parseJsonOrXml"
	"task1"
)

func Task1Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	if d.T1.Height <= 0 || d.T1.Width <= 0 {
		reason = "Sides must be integers"
		ok = false
	}
	return
}
func Task2Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task3Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task4Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task5Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task6Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task7Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}

func RiseAndShine() {
	Data := parse.GetData("data.json")
	if ok, reason := Task1Rules(Data); ok {
		//dowork
		fmt.Println(task1.DoTask1(Data.T1.Width, Data.T1.Height, Data.T1.Symbol))
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task2Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task3Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task4Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task5Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task6Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
	if ok, reason := Task7Rules(Data); ok {
		//dowork
	} else {
		//dontwork
		fmt.Println(ok, reason)
	}
}
