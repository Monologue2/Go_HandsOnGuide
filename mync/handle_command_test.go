package main

import (
	"bytes"
	"testing"
)

// handleCommand(w io.Writer, args []string) (err error)
func TestHandleCommand(t *testing.T) {
	usageMessage := `Usage: mync [http|grpc] -h

http: A HTTP Client.

http: <options> server

Options :
  -verb string
    	HTTP Method (default "GET")


grpc: A gRPC client.

grpc: <options> server


Options :
  -body string
    	Body of request
  -method string
    	Method to call
`

	testConfig := []struct {
		args   []string
		output string
		err    error
	}{
		{
			// 애플리케이션에 인수를 전달하지 않았을 경우
			args:   []string{},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		},
		{
			// -h 인수를 전달할 경우
			args:   []string{"-h"},
			err:    nil,
			output: usageMessage,
		},
		{
			// 알 수 없는 서브 커맨드가 지정된 경우
			args:   []string{"foo"},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		},
	}

	byteBuf := new(bytes.Buffer)

	for _, tc := range testConfig {
		err := handleCommand(byteBuf, tc.args)

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error %v, got %v", tc.err, err)
		}

		if len(tc.output) != 0 {
			gotOutput := byteBuf.String()
			if tc.output != gotOutput {
				t.Errorf("Expected output to be: %#v, Got: %#v", tc.output, gotOutput)
			}
		}
		byteBuf.Reset()
	}
}
