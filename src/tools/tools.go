package tools

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
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

func SplitOutput(i int) {
	fmt.Println("\n**************************")
	fmt.Printf("\n>>>>>>>>> Task %d <<<<<<<<<\n", i)
	fmt.Println("**************************\n")
}

func getExtension(fileName string) (string, error) {
	if ok, err := filepath.Match("*.js*", fileName); ok {
		return "json", err
	} else if ok, err := filepath.Match("*.xml*", fileName); ok {
		return "xml", err
	}
	return "", errors.New("Неправильный формат файла!\n")
}

func readFile(fileName string) ([]byte, error) {
	contents, err := ioutil.ReadFile(fileName)
	return contents, err
}

func makeStruct(D Data, contents []byte, extension string) (Data, error) {
	var err error

	if extension == "json" {
		err = json.Unmarshal(contents, &D)
	} else if extension == "xml" {
		err = xml.Unmarshal(contents, &D)
	}
	return D, err
}

func GetData(fileName string) (Data, error) {
	var MyData Data

	extension, err := getExtension(fileName)

	if err != nil {
		return MyData, err
	}

	contents, err := readFile(fileName)

	if err != nil {
		return MyData, err
	}

	MyData, err = makeStruct(MyData, contents, extension)

	return MyData, err
}
