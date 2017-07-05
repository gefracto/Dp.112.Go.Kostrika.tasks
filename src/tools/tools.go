package tools

import (
	"errors"
	"fmt"
)

var JsonRunners map[int]func([]byte) (error, []byte) = make(map[int]func([]byte) (error, []byte))

func RememberMe(task int, jsonRunner func([]byte) (error, []byte)) {
	JsonRunners[task] = jsonRunner
}

func RunTask(task int, jsonParams []byte) (error, []byte) {
	taskFunc, ok := JsonRunners[task]
	if !ok {
		return errors.New(fmt.Sprintf("Task #1 didn't registered runner")), nil
	}
	return taskFunc(jsonParams)
}
