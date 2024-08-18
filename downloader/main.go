package main

import (
	"fmt"
	"os"

	"github.com/Monologue2/Go_HandsOnGuide/downloader/downloader"
)

var ()

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stdout, "Must specify a HTTP URL to get data from")
		os.Exit(1)
	}

	body, err := downloader.FetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)
}
