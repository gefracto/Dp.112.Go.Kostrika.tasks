package tools

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task1"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task2"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task3"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task4"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task5"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task6"
	"github.com/gefracto/Dp.112.Go.Kostrika.tasks/task7"
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
	fmt.Println("\n\n**************************")
	fmt.Printf("********* Task %d *********\n", i)
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

func makeStruct(D Data, contents []byte, extension string) (error, Data) {
	var err error

	if extension == "json" {
		err = json.Unmarshal(contents, &D)
	} else if extension == "xml" {
		err = xml.Unmarshal(contents, &D)
	}
	return err, D
}

func GetData(fileName string) (error, Data) {
	var MyData Data

	extension, err := getExtension(fileName)

	if err != nil {
		return err, MyData
	}

	contents, err := readFile(fileName)

	if err != nil {
		return err, MyData
	}

	err, MyData = makeStruct(MyData, contents, extension)

	return err, MyData
}
