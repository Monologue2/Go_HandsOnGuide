package regpackage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// type FileHeader struct {
// 	Filename string
// 	Header   textproto.MIMEHeader
// 	Size     int64
// }

func packageRegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 결과를 놓아둘 구조체를 미리 메모리에 준비하고
		d := pkgRegisterResult{}
		err := r.ParseMultipartForm(5000) //메모리에 버퍼링 할 최대 바이트 수, 5kb
		if err != nil {
			http.Error(
				w, err.Error(), http.StatusBadRequest,
			)
			return
		}
		mForm := r.MultipartForm

		// multipart의 Form 구조체의 형식
		// type Form struct {
		// 	Value map[string][]string
		// 	File  map[string][]*FileHeader
		// }
		f := mForm.File["filedata"][0]
		d.Id = fmt.Sprintf(
			"%s-%s", mForm.Value["name"][0], mForm.Value["version"][0],
		)
		d.Filename = f.Filename
		d.Size = f.Size
		jsonData, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(
			w, "Invalid HTTP method specified.",
			http.StatusMethodNotAllowed,
		)
		return
	}
}

func startTestingServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(packageRegHandler))
}

func TestRegisterPackageData(t *testing.T) {
	ts := startTestingServer()
	defer ts.Close()

	p := multipartPkgData{
		Name:     "mypackage",
		Version:  "69.69",
		Filename: "mypackage.tar.gz",
		Bytes:    strings.NewReader("some_data"),
	}

	pResult, err := registerPackageData(ts.Client(), ts.URL, p)
	if err != nil {
		t.Fatal(err)
	}
	if pResult.Id != fmt.Sprintf("%s-%s", p.Name, p.Version) {
		t.Errorf("Expected package id to be %s-%s, Got %s", p.Name, p.Version, pResult.Id)
	}
	if pResult.Filename != p.Filename {
		t.Errorf("Expected package name to be %s, Got %s", p.Filename, pResult.Filename)
		if pResult.Size != 9 {
			t.Errorf("Expected package size to be 9, Got %d", pResult.Size)
		}
	}
}
