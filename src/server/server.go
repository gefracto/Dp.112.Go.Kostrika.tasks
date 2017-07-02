package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gefracto/qwer/src/tools"
)

func HandlerT1(w http.ResponseWriter, r *http.Request) {
	var d tools.Data
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	w.Write([]byte("Task 1\n"))
	if err, res := d.Dotask1(); err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		w.Write([]byte(res.(string)))
	}
}

func HandlerT2(w http.ResponseWriter, r *http.Request) {
	var d tools.Data
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	w.Write([]byte("\nTask 2\n"))
	if err, res := d.Dotask2(); err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		w.Write([]byte(strconv.Itoa(res.(int))))
	}
}
