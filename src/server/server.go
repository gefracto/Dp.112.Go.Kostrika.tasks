package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	t1 "github.com/gefracto/kostrika-go-tasks/src/task1"
	t2 "github.com/gefracto/kostrika-go-tasks/src/task2"
	"github.com/gefracto/kostrika-go-tasks/src/tools"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	addr, _ := strconv.Atoi(r.URL.Path[len("/task/"):])
	var d tools.Data
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	//	tmpl := ioutil.ReadFile("index.html")
	w.Write([]byte(fmt.Sprintf("Task %d\n", addr)))

	var err error
	var res string
	switch addr {

	case 1:
		T1 := d.T1
		err, res = t1.Dotask(T1.Width, T1.Height, T1.Symbol)

	case 2:

		err, res = t2.Dotask(d.T2.E1, d.T2.E2)

	case 3:

		if err, res := d.Dotask3(); err != nil {
			w.Write([]byte(error.Error(err)))
		} else {
			for _, v := range res {
				w.Write([]byte(v + " "))
			}

		}

	}
	if err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		w.Write([]byte(res))
	}

}

//func HandleAll(w http.ResponseWriter, r *http.Request) {
//	var d tools.Data
//	body, _ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(body, &d)

//	w.Write([]byte(res.(string)))

//}

//func HandlerT1(w http.ResponseWriter, r *http.Request) {
//	var d tools.Data
//	body, _ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(body, &d)
//	w.Write([]byte("Task 1\n"))
//	if err, res := d.Dotask1(); err != nil {
//		w.Write([]byte(error.Error(err)))
//	} else {
//		w.Write([]byte(res.(string)))
//	}
//}

//func HandlerT2(w http.ResponseWriter, r *http.Request) {
//	var d tools.Data
//	body, _ := iow.Write([]byte("\nTask 3\n"))
//	if err, res := d.Dotask3(); err != nil {
//		w.Write([]byte(error.Error(err)))
//	} else {
//		for _, v := range res {
//			w.Write([]byte(v + " "))
//		}

//	}util.ReadAll(r.Body)
//	json.Unmarshal(body, &d)
//	w.Write([]byte("\nTask 2\n"))
//	if err, res := d.Dotask2(); err != nil {
//		w.Write([]byte(error.Error(err)))
//	} else {
//		w.Write([]byte(strconv.Itoa(res.(int))))
//	}
//}

func HandlerT3(w http.ResponseWriter, r *http.Request) {
	var d tools.Data
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	w.Write([]byte("\nTask 3\n"))
	if err, res := d.Dotask3(); err != nil {
		w.Write([]byte(error.Error(err)))
	} else {
		for _, v := range res {
			w.Write([]byte(v + " "))
		}

	}
}
