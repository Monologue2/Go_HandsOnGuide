package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "20").Run(); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
	} // signal: killed
}
