package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gefracto/qwer/src/tools"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {

	addr := r.URL.Path[:]
	var d tools.Data
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	//	tmpl := ioutil.ReadFile("index.html")

	switch addr {

	case "/task/1":

		w.Write([]byte("Task 1\n"))
		if err, res := d.Dotask1(); err != nil {
			w.Write([]byte(error.Error(err)))
		} else {
			w.Write([]byte(res.(string)))
		}

	case "/task/2":

		w.Write([]byte("\nTask 2\n"))
		if err, res := d.Dotask2(); err != nil {
			w.Write([]byte(error.Error(err)))
		} else {
			w.Write([]byte(fmt.Sprintf("%d", res.(int))))
		}

	case "3":

		w.Write([]byte("\nTask 3\n"))
		if err, res := d.Dotask3(); err != nil {
			w.Write([]byte(error.Error(err)))
		} else {
			for _, v := range res {
				w.Write([]byte(v + " "))
			}

		}
	}

}

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
