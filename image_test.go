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
				Output:   "Hello World\n",
				ExitCode: 0,
			},
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

	}
}
