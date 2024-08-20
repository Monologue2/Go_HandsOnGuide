package regpackage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func packageTestServerHandler(w http.ResponseWriter, r *http.Request) {
	// Data 형식이 잘못 될 경우 Content-Type은 text/plain이 된다.
	// comma를 주의하세요
	data := `[
		{"name":"package1", "version":"1.1"},
		{"name":"package2", "version":"2.1"}
]`

	var packages []pkgData
	err := json.Unmarshal([]byte(data), &packages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(packages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func startTestPackageServer() *httptest.Server {
	// http.HandlerFunc 과 http.HandleFunc 의 차이는 무엇인가
	// http.HandlerFunc은 w ResponseWriter, r http.Request를 인자로 가지는
	// 함수를 http.Handle 로 변환시켜준다.
	// Http.HandleFunc은 http.Handle로 변환시킨 뒤 핸들러로 등록까지 함께 한다.
	// Mux(멀티플렉서)와 함께 사용한다.
	ts := httptest.NewServer(http.HandlerFunc(packageTestServerHandler))
	return ts
}

func TestFetchPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()

	packages, err := fetchPackageData(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, packages)
}
