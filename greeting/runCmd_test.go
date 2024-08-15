package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// type tests struct {
// 	// r io.Reader
// 	// w io.Writer
// 	c      config
// 	input  string
// 	output string
// 	err    error
// }

// func runCmd(r io.Reader, w io.Writer, c config) error
func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string // 테스트용 사용자 입력
		output string // 출력 결과
		err    error
	}{
		// {
		// 	c:      config{printUsage: true},
		// 	output: usageString,
		// },
		{
			c:      config{numTimes: 10},
			input:  "",
			output: strings.Repeat("Your name please? Press Enter key when done.\n", 1),
			err:    errors.New("You didn't enter your name"),
		},
		{
			c:      config{numTimes: 5},
			input:  "Leorca",
			output: "Your name please? Press Enter key when done.\n" + strings.Repeat("Nice to meet you Leorca\n", 5),
		},
	}

	// 사용자로부터의 입출력 동작을 따라해야한다.
	bytebuf := new(bytes.Buffer)
	for _, tc := range tests {
		rd := strings.NewReader(tc.input)
		err := runCmd(rd, bytebuf, tc.c)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got:%v", err)
		}

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error: %v, got error: %v", tc.err.Error(), err.Error())
		}

		gotMsg := bytebuf.String()
		if gotMsg != tc.output {
			t.Errorf("Expected stdout message: %v, got stdout message: %v", tc.output, gotMsg)
		}
		bytebuf.Reset()
	}
}
