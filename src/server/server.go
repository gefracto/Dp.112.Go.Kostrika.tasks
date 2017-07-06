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

	i, _ := strconv.Atoi(path.Base(r.URL.Path[:]))
	body, _ := ioutil.ReadAll(r.Body)
	_, res := tools.Run(i, body)
	w.Write(res)

}

// web.router
// restapi

//GET /COLLECTION/filter
// {...}
// GET /COLLECTION
// [{..}, {..}]
