package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func TestHelloHandler(t *testing.T) {
	const url = "http://localhost:8080/hello"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	body, _ := io.ReadAll(recorder.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
