package main

import (
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
	// "github.com/gefracto/kostrika-go-tasks/src/tools"
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
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.IndexHandler(w, r)
	})

	http.HandleFunc("/task/", server.HandleTask)
	http.ListenAndServe(":1111", nil)
}
