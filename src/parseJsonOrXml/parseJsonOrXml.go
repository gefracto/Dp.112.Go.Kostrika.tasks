package parseJsonOrXml

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"
	t2 "task2"
	t3 "task3"
	"fmt"
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

func GetData(fileName string) Data {
	var MyData Data
	extension := strings.Split(fileName, ".")
	if extension[len(extension)-1] == "json" {
		contents, _ := ioutil.ReadFile("data.json")
		json.Unmarshal(contents, &MyData)
	} else if extension[len(extension)-1] == "xml" {
		contents, _ := ioutil.ReadFile("data.xml")
		xml.Unmarshal(contents, &MyData)
	} else {
		fmt.Println("Не удалось открыть файл. \nРасширение файла должно быть формата json или xml.")
		return MyData
	}
	return MyData
}