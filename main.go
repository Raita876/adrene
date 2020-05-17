package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/math/fixed"
)

const (
	DST_DIR = "./dst"
)

type ImgMaker struct {
	width       int
	height      int
	marginTop   int
	marginLeft  int
	marginRight int
	fontSize    int
	lineSpace   int
}

func black() color.RGBA { return color.RGBA{0, 0, 0, 255} }

func white() color.RGBA { return color.RGBA{255, 255, 255, 255} }

func (im *ImgMaker) background() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, im.width, im.height))

	for i := 0; i < im.height; i++ {
		for j := 0; j < im.width; j++ {
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
		Size:              float64(im.fontSize),
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
		dr.Dot.X = fixed.I(im.marginLeft)
		dr.Dot.Y = fixed.I(im.marginTop + (im.fontSize+im.lineSpace)*i)
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
		if im.width-im.marginLeft-im.marginRight < dr.MeasureString(s[point:i+1]).Ceil() {
			sl = append(sl, s[point:i])
			point = i
		}
	}

	sl = append(sl, s[point:])

	return sl
}

func mkdir(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	r, err := Exec("docker --help")
	if err != nil {
		log.Fatal(err)
	}

	dstDir := DST_DIR

	err = mkdir(dstDir)
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("20160102150405")
	imgPath := fmt.Sprintf("./%s/%s.png", dstDir, date)
	text := r.Output

	im := ImgMaker{
		width:       800,
		height:      1600,
		marginTop:   40,
		marginLeft:  40,
		marginRight: 40,
		fontSize:    16,
		lineSpace:   4,
	}
	err = im.Create(imgPath, text)
	if err != nil {
		log.Fatal(err)
	}

}
