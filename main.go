package main

import (
	"github.com/gefracto/kostrika-go-tasks/src/server"
	"net/http"
)

func main() {

	http.HandleFunc("/task/", server.HandleTask)
	http.ListenAndServe(":1111", nil)
}
