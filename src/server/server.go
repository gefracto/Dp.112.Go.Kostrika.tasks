package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gefracto/kostrika-go-tasks/src/tools"
)

//func HandleAllTasks(w http.ResponseWriter, r *http.Request) {
//	body, _ := ioutil.ReadAll(r.Body)
//	err, res := tools.RunAll(body)
//	w.Write(res)
//	w.Header()
//	fmt.Println(err)
//	fmt.Println(string(res))

//}

func HandleTask(w http.ResponseWriter, r *http.Request) {

	i, _ := strconv.Atoi(r.URL.Path[len("/task/"):])

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(i)
	fmt.Println(string(body))

	err, res := tools.Run(i, body)
	w.Write(res)
	w.Header()

	fmt.Println(err)
	fmt.Println(string(res))

}

// web.router
// restapi

//GET /COLLECTION/filter
// {...}
// GET /COLLECTION
// [{..}, {..}]
