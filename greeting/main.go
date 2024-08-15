// Command Line Application, 인사 할 횟수를 전달받고 그만큼 커맨드 라인에 인사
// 커맨드라인 인터페이스 구현, 사용자 입력 전달, 애플리케이션 테스트
// 입력 -> 값 검증 -> 작업 -> 작업 결과 반환
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes   int
	printUsage bool
}

var (
	usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])
)

// func printUsage(w io.Writer) {
// 	fmt.Fprintf(w, usageString)
// }

// 입력 값 검사, 0이 넘지 않는 경우 Error 반환
func validateArgs(c config) error {
	if !(c.numTimes > 0 || c.printUsage) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

// 입력 인자는 os.Args[1:]
func parseArgs(w io.Writer, args []string) (config, error) {
	// var numTimes int
	// var err error

	c := config{}
	// FlagSet 객체 생성 (Application_Name, Error_Processing_Method)
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)

	//
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		return c, errors.New("Prositional arguments specified")
	}

	return c, nil
	// flag package를 사용하지 않고 수동으로 구현
	// // 1개 초과 또는 0개의 인자가 올 경우 Error
	// if len(args) != 1 {
	// 	return c, errors.New("Invalid number of arguments")
	// }

	// // 1개의 인자밖에 존재하지 않으므로 args[0]만 확인하면 됨
	// if args[0] == "-h" || args[0] == "--help" {
	// 	c.printUsage = true
	// 	return c, nil
	// }

	// numTimes, err = strconv.Atoi(args[0])
	// if err != nil {
	// 	return c, err
	// }

	// c.numTimes = numTimes
	// return c, nil
}

// io 패키지의 Reader, Writer 인터페이스를 매개변수로 사용
// 이렇게 할 경우 Use Case 테스트 코드를 작성하기 편해진다.
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press Enter key when done.\n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}

	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

// Reader, Writer Interface, config
func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		// printUsage(w)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		// printUsage(os.Stdout)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		// printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}
