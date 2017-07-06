package main

import (
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
)

func main() {

	http.HandleFunc("/task/", server.HandleTask)
	http.ListenAndServe(":1111", nil)
}
