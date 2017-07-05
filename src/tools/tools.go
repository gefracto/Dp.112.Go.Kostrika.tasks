package tools

import (
	"errors"
	"fmt"
)

var TasksMap map[int]func([]byte) (error, []byte) = make(map[int]func([]byte) (error, []byte))

func RememberMe(task int, taskFoo func([]byte) (error, []byte)) {
	TasksMap[task] = taskFoo
}

func Run(task int, json []byte) (error, []byte) {
	taskFunc, ok := TasksMap[task]
	if !ok {
		return errors.New(fmt.Sprintf("Task #%d didn't sign in", task)), nil
	}
	return taskFunc(json)
}

//func RunAll(json []byte) (error, []byte) {

//}
