package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func createContextWithTimeout(d time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), d)
	return ctx, cancel
}

func setupSignalHandler(w io.Writer, cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1) // os.Signal 을 받는 Buffer 1개를 가진 채널, capacity가 1이다. 라고 표현..
	// signal package의 Notify 함수는 시스템 호출을 받을 채널을 지정할 수 있다.
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Signal 대기를 위한 Go routine, 입력이 올 때까지 대기한다.
	go func() {
		s := <-c
		fmt.Fprintf(w, "Got signal: %v\n", s)
		cancelFunc()
	}()

	// Signal 이 제공되지 않는 동안에는 지정된 context 만큼 애플리케이션이 정상적으로 실행된다.
}

func executeCommand(ctx context.Context, command string, arg string) error {
	return exec.CommandContext(ctx, command, arg).Run()
}

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := exec.CommandContext(ctx, "sleep", "20").Run(); err != nil {
	// 	fmt.Fprintf(os.Stdout, "%v\n", err)
	// } // signal: killed

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <command> <argument>\n", os.Args[0])
		os.Exit(1)
	}

	command := os.Args[1]
	arg := os.Args[2]

	cmdTimeout := 30 * time.Second
	ctx, cancel := createContextWithTimeout(cmdTimeout)
	defer cancel()

	setupSignalHandler(os.Stdout, cancel)

	err := executeCommand(ctx, command, arg)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	// $ go run main.go sleep 60
	// signal: killed
	// exit status 1
}
