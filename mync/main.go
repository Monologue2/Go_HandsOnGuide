// flag library를 통해 외부 입력 분류하기
// 기능별로 Module 내의 Package 분리하기 *독립적인 기능을 가진 패키지를 만드는 법
package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Monologue2/Go_HandsOnGuide/mync/cmd"
)

var (
	errInvalidSubCommand = errors.New("Invalid sub-command specified")
)

// FlagSet의 출력을 전달받기 위한 w io.Writer
// -h, help, Command를 각 핸들러에 전달하여 출력
func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

// FlagSet의 출력을 전달받기 위한 w io.Writer
// Handler 함수를 args에 따라 맞춰 호출한다.
func handleCommand(w io.Writer, args []string) (err error) {
	// 첫 번째 요소 값이 http|grpc 인 경우 특정 동작 수행
	// -h|help 인 경우 사용 방법 제공
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			// 성공시 nil
			err = cmd.HandleHttp(w, args)
		case "grpc":
			// 성공시 nil
			err = cmd.HandleGrpc(w, args)
		case "-h":
			printUsage(w)
		case "help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}

	if errors.Is(err, cmd.ErrNoServerSpecified) || errors.Is(err, errInvalidSubCommand) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
