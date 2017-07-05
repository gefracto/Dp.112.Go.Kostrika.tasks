package tools

import (
	//t1 "github.com/gefracto/kostrika-go-tasks/src/task1"
	// "fmt"
	"errors"
	"fmt"
	// "github.com/gefracto/kostrika-go-tasks/src/task2"
)

// type Data struct {
// 	Task1 struct {
// 		Width  int    `json:"width"`
// 		Height int    `json:"height"`
// 		Symbol string `json:"symbol"`
// 	}

// 	Task2 struct {
// 		E1, E2 task2.Envelope
// 	}

// 	//	task2 struct{ task2.T2 }
// }

var JsonRunners map[int]func([]byte) (error, []byte) = make(map[int]func([]byte) (error, []byte))

func RegisterJsonRunner(task int, jsonRunner func([]byte) (error, []byte)) {
	JsonRunners[task] = jsonRunner
}

func RunTask(task int, jsonParams []byte) (error, []byte) {
	taskFunc, ok := JsonRunners[task]
	if !ok {
		return errors.New(fmt.Sprintf("Task #1 didn't registered runner")), nil
	}
	return taskFunc(jsonParams)
}
