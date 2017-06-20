package validator

import (
	"fmt"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}


func Task1Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task2Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task3Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task4Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task5Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task6Rules(s string) (ok bool, reason string, file string) {
	ok = true
	return
}
func Task7Rules(s string) (ok bool, reason string) {
	ok = true
	return
}

func ValidateData(file string) (ok bool, reason string) {
	var newFile string
	for i := 0; i < len(file); i++ {
		if  string(file[i]) != "{" &&
			string(file[i]) != "}" &&
			string(file[i]) != "[" &&
			string(file[i]) != "]" &&
			string(file[i]) != "\n" &&
			string(file[i]) != "," &&
			string(file[i]) != ":" &&
			string(file[i]) != "\"" &&
			string(file[i]) != "\t" &&
			string(file[i]) != " " {newFile += string(file[i])}
	}
	if ok, reason, file := Task1Rules(newFile); ok {
		newFile = file
	} else {
		return false, reason
	}

	if ok, reason, file := Task2Rules(newFile); ok {
		newFile = file
	} else {
		return false, reason
	}
	return true, newFile

}


