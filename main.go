package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
	_ "github.com/gefracto/kostrika-go-tasks/src/task1"
	_ "github.com/gefracto/kostrika-go-tasks/src/task2"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "index", nil)
}

func main() {

	http.HandleFunc("/task/", server.HandleTask)
	//	http.HandleFunc("/tasks/all", server.HandleAllTasks)
	http.ListenAndServe(":1111", nil)
}
