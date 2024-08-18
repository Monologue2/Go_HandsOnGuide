package downloader

import (
	"os"
	"strings"
	"testing"
)

func TestTcpHttpGet(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	url := strings.Split(ts.URL, "//")

	line, err := TcpHttpGet(os.Stdout, url[1])
	if err != nil {
		t.Fatalf("Got err %v", err)
	}
	expected := "Hello world"
	if line != expected {
		t.Fatalf("Expected output:%v, Got: %v", "Hello world", line)
	}
}
