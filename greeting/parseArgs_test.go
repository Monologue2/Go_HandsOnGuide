package main

import (
	"bytes"
	"errors"
	"testing"
)

// func parseArgs(args []string) (config, error)
type testConfig struct {
	args     []string
	err      error
	numTimes int
	// config
}

//	type config struct {
//		numTimes   int
//		printUsage bool
//	}
func TestParseArgs(t *testing.T) {
	// TestCases
	tests := []testConfig{
		{
			args:     []string{"-h"},
			err:      errors.New("flag: help requested"),
			numTimes: 0,
			// config: config{numTimes: 0, printUsage: true},
		},
		{
			args:     []string{"-n", "10"},
			err:      nil,
			numTimes: 10,
			// config: config{numTimes: 10, printUsage: false},
		},
		{
			args:     []string{"-n", "abc"},
			numTimes: 0,
			// err:      errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			err: errors.New("invalid value \"abc\" for flag -n: parse error"),
			// config: config{numTimes: 0, printUsage: false},
		},
		{
			args: []string{"-n", "1", "foo"},
			// err:      errors.New("Invalid number of arguments"),
			err:      errors.New("Prositional arguments specified"),
			numTimes: 1,
			// config: config{numTimes: 0, printUsage: false},
		},
	}

	//io.Writer용  공백 Buffer 객체 byte.Buffer
	byteBuf := new(bytes.Buffer)
	// Test 시작
	for _, tc := range tests {
		c, err := parseArgs(byteBuf, tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be; %v, got: %v\n",
				tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}

		// if c.printUsage != tc.config.printUsage {
		// 	t.Errorf("Expected printUsage to be :%v, got: %v\n",
		// 		tc.config.printUsage, c.printUsage)
		// }

		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to %v, got: %v\n",
				tc.numTimes, c.numTimes)
		}
		byteBuf.Reset()
	}
}
