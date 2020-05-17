package main

import (
	"errors"
	"log"
	"os"

	"adrene/command"
	"adrene/image"

	"github.com/urfave/cli/v2"
)

var (
	version string
	name    string
)

const (
	DST_DIR = "./dst"
)

type Option interface {
	apply(*image.ImgMaker)
}

func Run(cmd []string, imgPath string, opts ...Option) error {
	im := &image.ImgMaker{
		Width:        800,
		LimitHeight:  2400,
		MarginTop:    40,
		MarginLeft:   40,
		MarginRight:  40,
		MarginBottom: 0,
		FontSize:     16,
		LineSpace:    0,
		FontType:     "gomonobold",
	}

	for _, o := range opts {
		o.apply(im)
	}

	r, err := command.Exec(cmd...)
	if err != nil {
		return err
	}

	text := r.Output

	err = im.Create(imgPath, text) //TODO: change args, text -> command.Result
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Version: version,
		Name:    name,
		Usage:   "Adrene is a cli tool that can save the command execution result locally in png format.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "image",
				Aliases: []string{"i"},
				Value:   "out.png",
				Usage:   "Output image path.",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return errors.New("Argument is not set.")
			}

			imgPath := c.String("image")
			cmd := c.Args().Slice()
			err := Run(cmd, imgPath)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
