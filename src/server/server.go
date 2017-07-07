package server

import (
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	_ "github.com/gefracto/kostrika-go-tasks/src/task1"
	_ "github.com/gefracto/kostrika-go-tasks/src/task2"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	if i, err := strconv.Atoi(path.Base(r.URL.Path[:])); err != nil {
		w.Write(errors.New("URL не распознан"))
	}

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		w.Write(errors.New("HTTP не прочитан"))
	}

	if err, res := tools.Run(i, body); err != nil {
		w.Write(error.Error(err))
	} else {
		w.Write(res)
	}

}
