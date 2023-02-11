package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	// Open the file which contains the string of a PID
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	// Pretend to read the pid and convert it to an int
	// Value is thrown away, since there's no real process
	// to kill
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate killing the process
	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file - %s", err)
	}

	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
