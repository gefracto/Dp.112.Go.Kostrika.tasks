package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gefracto/kostrika-go-tasks/src/server"
	_ "github.com/gefracto/kostrika-go-tasks/src/task1"
	_ "github.com/gefracto/kostrika-go-tasks/src/task2"
)

func main() {
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/task/", server.HandleTask)
	http.HandleFunc("/tasks", server.HandleTasks)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
