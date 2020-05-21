package command

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExec(t *testing.T) {
	tests := []struct {
		cmd  []string
		want Result
	}{
		{
			cmd: []string{"echo", "Hello World"},
			want: Result{
				Command:  []string{"echo", "Hello World"},
				Output:   "Hello World\n",
				ExitCode: 0,
			},
		},
		{
			cmd: []string{"ls", "xxxxxx"},
			want: Result{
				Command:  []string{"ls", "xxxxxx"},
				Output:   "ls: xxxxxx: No such file or directory\n",
				ExitCode: 1,
			},
		},
	}

	for _, tt := range tests {
		got, err := Exec(tt.cmd...)
		if err != nil {
			t.Error(err)
		}

		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Errorf("Result missmatch (-got +want):\n%s", diff)
		}
	}
}
