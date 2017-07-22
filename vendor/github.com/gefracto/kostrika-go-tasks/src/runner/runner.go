package runner

import (
	"errors"
	"fmt"
	"github.com/gefracto/kostrika-go-tasks/src/tasklist"
)

func Run(task int, json []byte) (error, []byte) {
	taskFunc, ok := tasklist.TasksMap[task]
	if !ok {
		return errors.New(fmt.Sprintf("Таск #%d не зарегистрирован", task)), nil
	}
	return taskFunc(json)
}
