package parseJsonOrXml

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"
	t2 "task2"
	t3 "task3"
)

type Data struct {
	T1 struct {
		Width, Height int
		Symbol        string
	}
	T2 struct {
		Env1, Env2 t2.Envelope
	}
	T3 struct {
		SliceOfTriangles []t3.Triangle
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

func simpleSpellCheckerJson(file string) (b bool, message string) {
	openingCurlyBraces := strings.Count(file, "{")
	closingCurlyBraces := strings.Count(file, "}")
	quotes := strings.Count(file, "\"")
	openingBrackets := strings.Count(file, "[")
	closingBrackets := strings.Count(file, "]")

	if (openingBrackets + closingBrackets)%2 != 0 {
		b = false
		message += "К-ство открывающих и закрывающих квадратных скобок не равно\n"
	} else if (openingCurlyBraces + closingCurlyBraces)%2 != 0 {
		b = false
		message += "К-ство открывающих и закрывающих фигурных скобок не равно\n"
	} else if quotes%2 != 0 {
		b = false
		message += "К-ство кавычек не парно\n"
	} else {
		b = true
	}
	return b, message
}

func simpleSpellCheckerXml(file string) (b bool, message string) {
	openingAngleBrackets := strings.Count(file, "<")
	closingAngleBrackets := strings.Count(file, ">")
	slashes := strings.Count(file, "/")

	if openingAngleBrackets != closingAngleBrackets {
		b = false
		message += "К-ство открывающих и закрывающих угловых скобок не равно\n"
	} else if (openingAngleBrackets/2 != slashes) || (closingAngleBrackets/2 != slashes) {

		b = false
		message += "К-ство слэшей не то\n"
	} else {
		b = true
	}
	return b, message
}

func checkForNames(file string) (b bool, message string) {
	b = true

	var rules []string = []string{
		"T1", "T2", "T3", "T4", "T5", "T6", "T7",
		"Width", "Height", "Symbol",
		"Env1", "Env2", "Side1", "Side2",
		"SliceOfTriangles", "Name", "A", "B", "C",
		"Number",
		"Min", "Max",
		"Length", "MaxSquare",
		"File"}

	for _, v := range rules {
		if strings.Contains(file, v) == false {
			b = false
			message += v + "  "
		}
	}
	if b == false {
		message += " \n\nОтсутствуют либо неправильно написаны перечисленные выше переменные"
	}
	return b, message
}

func GetData(fileName string) (Data, bool, string) {
	var MyData Data
	extension := strings.Split(fileName, ".")

	if extension[len(extension)-1] == "json" {
		contents, _ := ioutil.ReadFile("data.json")

		if b, m := simpleSpellCheckerJson(string(contents)); b == false {
			return MyData, false, m
		}

		if b, m := checkForNames(string(contents)); b == false {
			return MyData, false, m
		}

		json.Unmarshal(contents, &MyData)

	} else if extension[len(extension)-1] == "xml" {
		contents, _ := ioutil.ReadFile("data.xml")

		if b, m := simpleSpellCheckerXml(string(contents)); b == false {
			return MyData, false, m
		}

		if b, m := checkForNames(string(contents)); b == false {
			return MyData, false, m
		}

		xml.Unmarshal(contents, &MyData)

	} else {
		return MyData, false, "Не удалось открыть файл. \nРасширение файла должно быть формата json или xml."
	}
	return MyData, true, ""
}
