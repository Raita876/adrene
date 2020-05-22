package main

import "testing"

func TestCreate(t *testing.T) {
	tests := []struct {
		imgPath string
		result  Result
	}{
		{
			imgPath: "tmp.png",
			result: Result{
				Command:  []string{"echo", "Hello World"},
				Output:   "Hello World",
				ExitCode: 0,
			},
		},
	}

	for _, tt := range tests {

	}
}
