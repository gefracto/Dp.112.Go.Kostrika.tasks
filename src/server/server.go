package server

import (
	"fmt"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"html/template"
	"io/ioutil"
	"net/http"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var err error
	var res []byte

	err, res = tools.RunTask(2, body)

	if err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		w.Write(res)
	}

}
