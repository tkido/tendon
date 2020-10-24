package ui

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// FontSize is font size
type FontSize int

// FontType is font type
type FontType int

// fontManager manage font
type fontManager struct {
	Fonts map[FontType]fontData
}

// fontData is font data
type fontData struct {
	Font  *truetype.Font
	Faces map[FontSize]font.Face
}

func (fm *fontManager) face(fontType FontType, size FontSize) font.Face {
	fd, ok := fm.Fonts[fontType]
	if !ok {
		log.Fatalf("fontManager.getFace: unknown FontType %d", fontType)
	}
	if face, ok := fd.Faces[size]; ok {
		return face
	}
	face := truetype.NewFace(fd.Font, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	fd.Faces[size] = face
	return face
}

// RegisterFont register font
func RegisterFont(fontType FontType, f *truetype.Font) {
	gm.Fonts[fontType] = fontData{f, map[FontSize]font.Face{}}
}
