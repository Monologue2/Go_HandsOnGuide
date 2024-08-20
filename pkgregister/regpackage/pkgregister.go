package regpackage

import (
	"encoding/json"
	"io"
	"net/http"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// 다음을 반환하는 API가 있다. 이를 처리하기
// [
// 	{"name":"package1", "version":"1.1"},
// 	{"name":"package2", "version":"2.1"},
// ]

func fetchPackageData(url string) ([]pkgData, error) {
	var packages []pkgData
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		return packages, err
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &packages) // Slice에 집어넣어도 잘 작동합니다. 세상에!
	return packages, err
}
