package main

import (
	"io/ioutil"
	"log"

	"github.com/tkido/tendon/assets"
	"github.com/tkido/tendon/ui"

	"github.com/golang/freetype/truetype"
)

// Font const
const (
	FontRegular ui.FontType = iota
	FontBold
	FontPixel

	FontXS     ui.FontSize = 12
	FontSmall              = 16
	FontMedium             = 24
	FontLarge              = 36
	FontXL                 = 48
)

func init() {
	registerFont(FontRegular, "mplus-1p-regular.ttf")
	registerFont(FontBold, "mplus-1p-bold.ttf")
	registerFont(FontPixel, "PixelMplus12-Regular.ttf")
}

func registerFont(fontType ui.FontType, path string) {
	ttf, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer ttf.Close()
	bs, err := ioutil.ReadAll(ttf)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(bs)
	if err != nil {
		log.Fatal(err)
	}
	ui.RegisterFont(fontType, tt)
}
