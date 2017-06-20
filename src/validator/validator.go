package validator

import (
	"fmt"
	"strings"
	"strconv"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
//T1Width16Height8Symbol*T2Env1Side121Side215Env2Side13.5Side222.3T3SliceOfTrianglesNameABCA10.3B10.1C10NameDEFA10.1B10C10.3NameGHIA34B43.2C20NameJKLA2B3C2T4Number41212226T5Min999990Max999999T6Length10MaxSquare25T7Filecontext

func Task1Rules(str string) (ok bool, reason string, file string) {
	ok = true

	leftWidth := strings.Index(str, "Width") + len("Width")
	rightWidth := strings.Index(str, "Height")

	if _, err := strconv.Atoi(string(str[leftWidth:rightWidth])); err != nil {
		ok = false
		reason = "\"Width\" должно быть числом типа int и без кавычек"
	}
	//fmt.Println(strconv.Atoi(string(str[leftWidth:rightWidth])))

	leftHeight := strings.Index(str, "Height") + 6
	rightHeight := strings.Index(str, "Symbol")

	if _, err := strconv.Atoi(string(str[leftHeight:rightHeight])); err != nil {
		ok = false
		reason = "\"Height\" должно быть числом типа int и без кавычек"
	}
	//fmt.Println(strconv.Atoi(string(str[leftHeight:rightHeight])))
	// Symbol не проверяем
	beginFileWithIndex := strings.Index(str, "Env1Side1") + len("Env1Side1")
	file = string(str[beginFileWithIndex:])
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
	//var newFile string

	if ok, reason, file = Task1Rules(file); ok == false {
		return false, reason
	}

	if ok, reason, file = Task2Rules(file); ok == false {
		return false, reason
	}

	if ok, reason, file = Task3Rules(file); ok == false {
		return false, reason
	}

	if ok, reason, file = Task4Rules(file); ok == false {
		return false, reason
	}

	if ok, reason, file = Task5Rules(file); ok == false {
		return false, reason
	}

	if ok, reason, file = Task6Rules(file); ok == false {
		return false, reason
	}

	if ok, reason = Task7Rules(file); ok == false {
		return false, reason
	}

	return true, reason

}
