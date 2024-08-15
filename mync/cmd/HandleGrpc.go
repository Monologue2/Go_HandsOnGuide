package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	server string
	method string
	body   string
}

// 기능에 따라 Flag 를 분리시킨다 = 저장할 변수의 위치와 동작이 달라진다.
func HandleGrpc(w io.Writer, args []string) (err error) {
	c := grpcConfig{}
	// FlagSet 선언, 이름과 오류 발생시 대처
	fs := flag.NewFlagSet("HandleGrpc", flag.ContinueOnError)
	// 오류등의 출력을 위한 io.Writer 인터페이스 제공
	fs.SetOutput(w)
	// Parameter 조정
	fs.StringVar(&c.method, "method", "", "Method to call")
	fs.StringVar(&c.body, "body", "", "Body of request")

	// flagSet Interface의 사용법
	fs.Usage = func() {
		var usageString = `

grpc: A gRPC client.

grpc: <options> server
`
		fmt.Fprintf(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options :")
		fs.PrintDefaults()
	}

	err = fs.Parse(args)
	if err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	c.server = fs.Arg(0)
	fmt.Fprintln(w, "Excuting grpc command")

	return nil
}
