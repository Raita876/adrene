package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"sync"
)

// Result is a structure that holds the execution result of the command
type Result struct {
	Command  []string
	Output   string
	ExitCode int
}

// Exec returns the execution result of the passed command.
func Exec(cmd ...string) (Result, error) {
	var c *exec.Cmd

	if len(cmd) < 2 {
		c = exec.Command(cmd[0])
	} else {
		c = exec.Command(cmd[0], cmd[1:]...)
	}

	rStdout, err := c.StdoutPipe()
	if err != nil {
		return Result{}, err
	}

	rStderr, err := c.StderrPipe()
	if err != nil {
		return Result{}, err
	}

	out := bytes.NewBuffer(nil)

	wStdout := io.MultiWriter(out, os.Stdout)
	wStderr := io.MultiWriter(out, os.Stderr)

	err = c.Start()
	if err != nil {
		return Result{}, err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		io.Copy(wStdout, rStdout)
		rStdout.Close()
		wg.Done()
	}()

	go func() {
		io.Copy(wStderr, rStderr)
		rStderr.Close()
		wg.Done()
	}()

	wg.Wait()

	err = c.Wait()
	if _, ok := err.(*exec.ExitError); err != nil && !ok {
		return Result{}, err
	}

	r := Result{
		Command:  cmd,
		Output:   string(out.Bytes()),
		ExitCode: c.ProcessState.ExitCode(),
	}

	return r, nil
}
