package server

import (
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
}
func Test_HandleTasks(t *testing.T) {

}
func Test_HandleTask(t *testing.T) {

}
