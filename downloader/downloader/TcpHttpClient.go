package downloader

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func TcpHttpGet(w io.Writer, url string) (string, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	request := "GET / HTTP/1.1\r\n" +
		"Host: " + url + "\r\n" +
		"Connection: close\r\n" +
		"\r\n"

	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", err
	}

	response := bufio.NewReader(conn)
	contentLength := 0

	fmt.Print("/// This is Header ///\n") // This is header
	for {
		line, err := response.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if line == "\r\n" {
			break
		}

		if strings.HasPrefix(line, "Content-Length:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				contentLength, err = strconv.Atoi(strings.TrimSpace(parts[1]))
				if err != nil {
					log.Fatal("Invalid Content-Length")
				}
			}
		}
		fmt.Print(line) // This is header
	}
	body := make([]byte, contentLength) // contentLength 크기를 가진 []byte slice 생성
	// Http 패키지 내에서 io.ReadFull 을 쓸 가능성 매우 높다.
	// io.ReadAll 과 io.ReadFull의 차이
	// io.ReadAll: io.Reader에서 모든 데이터를 읽은 뒤 []byte Slice 반환
	// io.ReadFull: io.Reader에서 데이터를 읽되, 정해진 buf만큼 읽음
	_, err = io.ReadFull(response, body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), nil
}
