package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	totalDuration time.Duration = 5
)

func getName(r io.Reader, w io.Writer) (name string, err error) {
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done."
	fmt.Fprintln(w, msg)
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		return "", err
	}
	name = scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You entered an empty name.")
	}
	return name, nil
}

func getNameContext(ctx context.Context, r io.Reader, w io.Writer) (name string, err error) {
	name = "Default Name"

	c := make(chan error, 1) // Error Interface를 전달하는 channal 생성, buffer size는 1

	go func() {
		name, err = getName(r, w)
		c <- err // err를 채널로 전송
	}()

	select {
	case <-ctx.Done(): // 5초 경과로 인해 Context가 종료된 err
		return name, ctx.Err()
	case err = <-c: // nil 또는 사용자 입력 상 에러로 인한 err
		return name, err
	}
}

func main() {
	allowedDuration := totalDuration * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), allowedDuration)
	// main 종료 전 반드시 context를 종료시키는 defer cancel()
	defer cancel()

	name, err := getNameContext(ctx, os.Stdin, os.Stdout)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, name)
}
