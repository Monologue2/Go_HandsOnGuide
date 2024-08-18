package downloader

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello world")
			}))
	return ts
}

// Go Syntax, Trailing comma: in composite literals or multiline function calls, the trailing comma is mandatory if the closing delimiter (\), \], or \}) is on a new line.
// func NewTestServer() *httptest.Server {
// 	ts := httptest.NewServer(
// 		http.HandlerFunc(
// 			func(w http.ResponseWriter, r *http.Request) {
// 				fmt.Fprintf(w, "Hello World")
// 			},
// 		),
// 	)
// 	return  ts
// }

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	data, err := FetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello world"
	if string(data) != expected {
		t.Errorf("Expected response data to be: %s, Got: %s", expected, data)
	}
}
