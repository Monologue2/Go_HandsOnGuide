package regpackage

// mime/multipart Package를 사용하고, 데이터를 캡슐화하는 방법 배우고
// 받은 데이터를 Marshal, Unmarshal 하여 직렬화, 역직렬화 하는 법을 배움

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type pkgRegisterResult struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// 비즈니스 로직, 데이터 생성
func registerPackageData(client *http.Client, url string, data multipartPkgData) (pkgRegisterResult, error) {
	p := pkgRegisterResult{}
	payload, contentType, err := createMultiPartMessage(data)
	if err != nil {
		return p, err
	}

	// payload는 []byte type

	reader := bytes.NewReader(payload)
	r, err := client.Post(url, contentType, reader)
	if err != nil {
		return p, nil
	}
	defer r.Body.Close()

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return p, nil
	}
	err = json.Unmarshal(respData, &p)
	return p, err
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	client := http.Client{Timeout: d}
	return &client
}

type multipartPkgData struct {
	Name     string
	Version  string
	Filename string
	Bytes    io.Reader
}

// 비즈니스 로직 부분, Data 패키징
func createMultiPartMessage(data multipartPkgData) ([]byte, string, error) {
	var b bytes.Buffer
	var err error
	var fw io.Writer

	mw := multipart.NewWriter(&b)

	//mime/multipart/form-data 로 인코딩, Buffer에 집어넣고 있으므로..,.
	fw, err = mw.CreateFormField("name")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprint(fw, data.Name)

	fw, err = mw.CreateFormField("version")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprint(fw, data.Version)

	fw, err = mw.CreateFormFile("filedata", data.Filename)
	if err != nil {
		return nil, "", err
	}
	// byte data를 Writer로 전송하는 방법?
	_, err = io.Copy(fw, data.Bytes)
	if err != nil {
		return nil, "", err
	}

	err = mw.Close()
	if err != nil {
		return nil, "", err
	}

	contentType := mw.FormDataContentType()
	return b.Bytes(), contentType, nil
}
