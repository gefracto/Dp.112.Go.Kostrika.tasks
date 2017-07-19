package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	"github.com/gefracto/kostrika-go-tasks/src/runner"
	//	"github.com/gefracto/kostrika-go-tasks/templates"
)

type Answer struct {
	Task   int
	Resp   string
	Reason string
}

func getData(a interface{}) []byte {
	data, _ := json.Marshal(a)
	return data
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	var Answers []Answer
	m := make(map[int]interface{})
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &m)

	for key, value := range m {
		var A Answer
		A.Task = key
		err, val := runner.Run(key, getData(value))
		A.Reason = fmt.Sprint(err)
		A.Resp = string(val)
		Answers = append(Answers, A)
	}
	js, _ := json.Marshal(Answers)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HandleTask(w http.ResponseWriter, r *http.Request) {

	i, err := strconv.Atoi(path.Base(r.URL.Path[:]))

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}
	var A Answer
	A.Task = i
	err, val := runner.Run(i, body)
	A.Reason = fmt.Sprint(err)
	A.Resp = string(val)
	js, _ := json.Marshal(A)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json.RawMessage(js))

}
