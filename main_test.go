package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	tests := []struct {
		cmd     []string
		imgPath string
		opts    []Option
		want    string
	}{
		{
			cmd:     []string{"echo", "Hello World"},
			imgPath: "tmp.png",
			opts:    []Option{},
			want:    "test/want.png",
		},
	}

	for _, tt := range tests {
		err := Run(tt.cmd, tt.imgPath, tt.opts...)
		if err != nil {
			t.Fatal(err)
		}

		got, err := FileBytes(tt.imgPath)
		if err != nil {
			t.Fatal(err)
		}

		want, err := FileBytes(tt.want)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("Result missmatch (-got +want):\n%s", diff)
		}
	}

}
