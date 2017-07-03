package main

import (
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
)

func main() {

	//	arg := flag.String("file", "data.json",
	//		"Usage: -file=fileName.extension")

	//	flag.Parse()

	//	_, data := tools.GetData(*arg)
	//	tools.Runall(data)

	http.HandleFunc("/task/1", server.HandlerT1)
	http.HandleFunc("/task/2", server.HandlerT2)
	http.HandleFunc("/task/3", server.HandlerT3)
	http.ListenAndServe(":1111", nil)
}
