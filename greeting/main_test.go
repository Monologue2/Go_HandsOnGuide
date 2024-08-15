package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestWithArgument(t *testing.T) {
	cmd := exec.Command("./application", "-h")

	err := cmd.Run()
	if err != nil {
		t.Fatalf("process ran with err%v, want exit status 0", err)
	}

	name := "Leorca"
	// StdinPipe() 함수로 애플리케이션 표준 입력에 접근
	cmd = exec.Command("./application", "5")
	stdin, err := cmd.StdinPipe()

	// 표준 출력 캡쳐
	output := &bytes.Buffer{}
	cmd.Stdout = output
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start command: %v", err)
	}

	// 표준 입력 데이터 전달
	if _, err := stdin.Write([]byte(name)); err != nil {
		t.Fatalf("Failed to write to stdin: %v", err)
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	expectedOutput := "Your name please? Press Enter key when done.\n" + strings.Repeat("Nice to meet you Leorca\n", 5)
	if output.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output.String())
	}
}
