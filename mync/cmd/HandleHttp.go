package cmd

import (
	"flag"
	"fmt"
	"io"
)

type httpConfig struct {
	url  string
	verb string // What is verb? 동사요?, 프로그래밍에서 어떤 의미를 갖는가
}

func HandleHttp(w io.Writer, args []string) (err error) {
	var v string

	// flag 선언 과정
	fs := flag.NewFlagSet("HandleHttp", flag.ContinueOnError)
	fs.SetOutput(w)
	// func (f *flag.FlagSet) StringVar(p *string, name string, value string, usage string)
	fs.StringVar(&v, "verb", "GET", "HTTP Method")

	fs.Usage = func() {
		var usageString = `
http: A HTTP Client.

http: <options> server`
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

	// url이 위치 인자(Positional arguments)로 전달되지 않았을 경우
	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	c := httpConfig{verb: v}
	// 0 번째 위치인자는 url, Arg(0)
	c.url = fs.Arg(0)
	fmt.Fprintln(w, "Excuting http command")
	return nil
}
