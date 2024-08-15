package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [cmd-a|cmd-b] -h\n", os.Args[0])
	handleCmda(w, []string{"-h"})
	handleCmdb(w, []string{"-h"})
}

func handleCmda(w io.Writer, args []string) (err error) {
	var v string
	// FlagSet 선언 -> SetOutput -> Set Variables(varibale_address, command, value, help)
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")

	err = fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Excuting command A\n")
	return nil
}

func handleCmdb(w io.Writer, args []string) (err error) {
	var v string
	// FlagSet 선언 -> SetOutput -> Set Variables(varibale_address, command, value, help)
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 2")

	err = fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Excuting command B\n")
	return nil
}

func main() {
	var err error
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(1)
	}

	// 호출하는 함수가 다르다. == 기능 별로 command를 분리했다.
	switch os.Args[1] {
	case "cmd-a":
		err = handleCmda(os.Stdout, os.Args[2:])
	case "cmd-b":
		err = handleCmdb(os.Stdout, os.Args[2:])
	default:
		printUsage(os.Stdout)
	}

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(1)
}
