package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/math/fixed"
)

type ImgMaker struct {
	Width       int
	Height      int
	MarginTop   int
	MarginLeft  int
	MarginRight int
	FontSize    int
	LineSpace   int
}

func black() color.RGBA { return color.RGBA{0, 0, 0, 255} }

func white() color.RGBA { return color.RGBA{255, 255, 255, 255} }

func (im *ImgMaker) background() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, im.Width, im.Height))

	for i := 0; i < im.Height; i++ {
		for j := 0; j < im.Width; j++ {
			img.Set(j, i, black())
		}
	}

	return img
}

func (im *ImgMaker) face() (font.Face, error) {
	var face font.Face

	ft, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		return face, err
	}

	opt := truetype.Options{
		Size:              float64(im.FontSize),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	face = truetype.NewFace(ft, &opt)

	return face, nil
}

func (im *ImgMaker) Create(imgPath string, text string) error {

	img := im.background()
	face, err := im.face()
	if err != nil {
		return err
	}

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(white()),
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	for i, s := range im.textToList(dr, text) {
		dr.Dot.X = fixed.I(im.MarginLeft)
		dr.Dot.Y = fixed.I(im.MarginTop + (im.FontSize+im.LineSpace)*i)
		dr.DrawString(s)
	}

	file, err := os.Create(imgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func (im *ImgMaker) textToList(dr *font.Drawer, text string) []string {
	var sl []string
	for _, s := range strings.Split(text, "\n") {
		sl = append(sl, im.stringToList(dr, s)...)
	}

	return sl
}

func (im *ImgMaker) stringToList(dr *font.Drawer, s string) []string {
	var sl []string

	point := 0
	for i := 1; i < len(s)-1; i++ {
		if im.Width-im.MarginLeft-im.MarginRight < dr.MeasureString(s[point:i+1]).Ceil() {
			sl = append(sl, s[point:i])
			point = i
		}
	}

	sl = append(sl, s[point:])

	return sl
}
