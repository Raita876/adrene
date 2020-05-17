package main

import "os/exec"

type Result struct {
	Command  string
	Output   string
	ExitCode int
}

func Exec(cmd string) (Result, error) {
	c := exec.Command("bash", "-co", "pipefail", cmd)

	out, err := c.CombinedOutput()
	if err != nil {
		return Result{}, err
	}

	r := Result{
		Command:  cmd,
		Output:   string(out),
		ExitCode: c.ProcessState.ExitCode(),
	}

	return r, nil
}
