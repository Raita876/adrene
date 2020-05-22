package main

import (
	"os"
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
			want: "want.png",
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

		got, err := fileBytes(tt.imgPath)
		if err != nil {
			t.Fatal("Failure get bytes png file")
		}

		want, err := fileBytes(tt.want)
		if err != nil {
			t.Fatal("Failure get bytes png file")
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("Result missmatch (-got +want):\n%s", diff)
		}

	}
}

func fileBytes(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	b := []byte{}
	_, err = f.Read(b)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
