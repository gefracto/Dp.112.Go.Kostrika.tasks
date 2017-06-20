package validator

import (
	"fmt"
	"strings"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func Task1Rules(s string) (ok bool, reason string, file string) {
	ok = true
	strings.Index(s,"T1")
	//strconv.Parse
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
func JsonToStr(file string) (newFile string) {
	for i := 0; i < len(file); i++ {
		if string(file[i]) != "{" &&
			string(file[i]) != "}" &&
			string(file[i]) != "[" &&
			string(file[i]) != "]" &&
			string(file[i]) != "\n" &&
			string(file[i]) != "," &&
			string(file[i]) != ":" &&
			string(file[i]) != "\"" &&
			string(file[i]) != "\t" &&
			string(file[i]) != " " {
			newFile += string(file[i])
		}
	}
	return newFile
}

func XmlToStr(file string) (newFile string) {
	for i := 0; i < len(file); i++ {
		if string(file[i]) != "<" &&
			string(file[i]) != ">" &&
			string(file[i]) != "/" &&
			string(file[i]) != "\n" &&
			string(file[i]) != "\t" {
			newFile += string(file[i])
		}
	}
	return newFile
}

func ValidateData(contents []byte, format string) (ok bool, reason string) {
	var file string
	for _, v := range contents {
		file += string(v)
	}
	if format == "json" {
		file = JsonToStr(file)
	} //else if format == "xml" {
		//file = XmlToStr(file)
	//}

	fmt.Println(file)
	//var newFile string

	return true, reason

}
