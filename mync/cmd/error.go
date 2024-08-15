package cmd

import (
	"errors"
)

var (
	ErrNoServerSpecified = errors.New("You have to specified the remote server.")
)
