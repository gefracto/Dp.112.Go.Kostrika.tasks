package server

import (
	"fmt"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	i, _ := strconv.Atoi(r.URL.Path[len("/task/task"):])

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(i)
	fmt.Println(string(body))

	var err error
	var res []byte

	err, res = tools.RunTask(i, body)

	if err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		w.Write(res)
	}

	fmt.Println(err)
	fmt.Println(string(res))

}
