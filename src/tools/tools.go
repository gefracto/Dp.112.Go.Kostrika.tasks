package tools

import (
	"errors"
	"fmt"
)

var TasksMap map[int]func([]byte) (error, []byte) = make(map[int]func([]byte) (error, []byte))

func RememberMe(task int, taskFoo func([]byte) (error, []byte)) {
	TasksMap[task] = taskFoo
}

func RunTask(task int, json []byte) (error, []byte) {
	taskFunc, ok := TasksMap[task]
	if !ok {
		return errors.New(fmt.Sprintf("Task #1 didn't sign in")), nil
	}
	return taskFunc(json)
}
