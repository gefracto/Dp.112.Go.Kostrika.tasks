package main

import (
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
	_ "github.com/gefracto/kostrika-go-tasks/src/task1"
	_ "github.com/gefracto/kostrika-go-tasks/src/task2"
)

func main() {

	http.HandleFunc("/task/", server.HandleTask)
	http.ListenAndServe(":1111", nil)
}
