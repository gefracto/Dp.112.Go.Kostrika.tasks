package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/gefracto/kostrika-go-tasks/src/runner"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {
	//	var i int
	//	var body []byte

	i, err := strconv.Atoi(path.Base(r.URL.Path[:]))

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}

	err, res := runner.Run(i, body)

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	} else {
		w.Write(res)
	}

}
