package server

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getData(t *testing.T) {

}
func Test_IndexHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := template.ParseFiles("templates/index.html")
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func Test_HandleTasks(t *testing.T) {

}
func Test_HandleTask(t *testing.T) {

}
