package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		imgPath string
		result  Result
		want    string
	}{
		{
			imgPath: "tmp.png",
			result: Result{
				Command:  []string{"echo", "Hello World"},
				Output:   "Hello World\n",
				ExitCode: 0,
			},
			want: "test/want.png",
		},
	}

	im := &ImgMaker{
		Width:        800,
		LimitHeight:  2400,
		MarginTop:    40,
		MarginLeft:   40,
		MarginRight:  40,
		MarginBottom: 0,
		FontSize:     16,
		LineSpace:    4,
		FontType:     "gomonobold",
		Prompt:       "$",
	}

	for _, tt := range tests {
		err := im.Create(tt.imgPath, tt.result)
		if err != nil {
			t.Error("Failure ImgMaker.Create()")
		}

		got, err := FileBytes(tt.imgPath)
		if err != nil {
			t.Fatal("Failure get bytes png file")
		}

		want, err := FileBytes(tt.want)
		if err != nil {
			t.Fatal("Failure get bytes png file")
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("Result missmatch (-got +want):\n%s", diff)
		}

	}
}
