package tools

import (
	"errors"
	"fmt"
)

type task func([]byte) (error, []byte)

var TasksMap map[int]task = make(map[int]task)

func RememberMe(tasknum int, t task) {
	TasksMap[tasknum] = t
	//	panic(t)
}

func Run(task int, json []byte) (error, []byte) {
	taskFunc, ok := TasksMap[task]
	if !ok {
		return errors.New(fmt.Sprintf("Task #%d didn't sign in", task)), nil
	}
	return taskFunc(json)
}
