package tools

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gefracto/kostrika-go-tasks/src/task1"
	"github.com/gefracto/kostrika-go-tasks/src/task2"
	"github.com/gefracto/kostrika-go-tasks/src/task3"
	"github.com/gefracto/kostrika-go-tasks/src/task4"
	"github.com/gefracto/kostrika-go-tasks/src/task5"
	"github.com/gefracto/kostrika-go-tasks/src/task6"
	"github.com/gefracto/kostrika-go-tasks/src/task7"
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

//func Runall(data Data) {

//	numtasks := reflect.ValueOf(&data).Elem().NumField()
//	item := reflect.ValueOf(&data)

//	for i := 1; i <= numtasks; i++ {
//		pattern := "<task%d.T%d Value>"

//		if item.Elem().Field(i-1).String() != fmt.Sprintf(pattern, i, i) {
//			numtasks += 1
//			continue
//		}

//		methodname := fmt.Sprintf("Dotask%d", i)
//		SplitOutput(i)
//		output := item.MethodByName(methodname).Call(nil)

//		Err := output[0]
//		Dat := output[1]

//		if reflect.Value.IsNil(Err) {
//			fmt.Println(Dat)

//		} else {
//			fmt.Println(Err)
//		}
//	}

//}

func SplitOutput(i int) (s string) {
	s += "\n\n**************************"
	s += fmt.Sprintf("********* Task %d *********\n", i)
	s += "**************************\n"
	return
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
