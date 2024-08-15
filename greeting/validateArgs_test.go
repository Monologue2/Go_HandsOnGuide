package main

import (
	"errors"
	"testing"
)

// func validateArgs(c config) error
func TestValidateArgs(t *testing.T) {
	// 익명 구조체로 Test Struct 선언
	tests := []struct {
		c   config
		err error
	}{
		{
			c:   config{},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	// Test 시작
	for _, tc := range tests {
		err := validateArgs(tc.c)
		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Errorf("Expected error to be %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	}
}
