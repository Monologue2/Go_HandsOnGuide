package main

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

//	type Reader interface {
//	    Read(p []byte) (n int, err error)
//	}

type delayedReader struct {
	delay time.Duration
	data  string
}

// io.Reader 인터페이스 구현
func (dr *delayedReader) Read(p []byte) (n int, err error) {
	time.Sleep(dr.delay) // Simulate delay
	// strings.NewReader io.Reader를 가져옴.
	return strings.NewReader(dr.data).Read(p)
}

func TestGetNameContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	r := &delayedReader{delay: 2 * time.Second, data: "Leorca"}
	w := new(bytes.Buffer)
	name, err := getNameContext(ctx, r, w)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context deadline exceeded error, got %v", err)
	}

	if name != "Default Name" {
		t.Errorf("expected default name, got %v", name)
	}
}
