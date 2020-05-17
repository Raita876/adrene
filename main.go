package main

import (
	"fmt"
	"log"

	"adrene/command"
	"adrene/image"
)

const (
	DST_DIR = "./dst"
)

type Option interface {
	apply(*image.ImgMaker)
}

func Run(cmd []string, imgPath string, opts ...Option) error {
	im := &image.ImgMaker{
		Width:       800,
		Height:      2000,
		MarginTop:   40,
		MarginLeft:  40,
		MarginRight: 40,
		FontSize:    16,
		LineSpace:   4,
	}

	for _, o := range opts {
		o.apply(im)
	}

	r, err := command.Exec(cmd...)
	if err != nil {
		return err
	}

	text := r.Output

	err = im.Create(imgPath, text)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	imgPath := fmt.Sprintf("./%s.png", "out")
	err := Run([]string{"docker", "--help"}, imgPath)
	if err != nil {
		log.Fatal(err)
	}
}
