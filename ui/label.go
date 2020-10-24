package ui

import (
	"fmt"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

// Align is Horizontal Align
type Align int

// Align
const (
	Left Align = iota
	Center
	Right
)

// Label is simple box
type Label struct {
	*Box
	text string
	FontType
	FontSize
	FontColor color.Color
	Align
}

// SetText for ui.Texter interface
func (l *Label) SetText(s string) {
	l.text = s
	l.dirtySelf()
}

// Text for ui.Texter interface
func (l *Label) Text() string {
	return l.text
}

// NewLabel makes new *Label
func NewLabel(w, h int, text string, fontType FontType, fontSize FontSize, align Align, color, bgColor color.Color) *Label {
	b := NewBox(w, h, bgColor)
	l := &Label{b, text, fontType, fontSize, color, align}
	l.Self = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *Label) Reflesh() {
	l.Box.Reflesh()
	face := gm.face(l.FontType, l.FontSize)
	w, h := l.Size()
	textWidth := font.MeasureString(face, l.text).Ceil()
	var x int
	switch l.Align {
	case Left:
		x = 0
	case Right:
		x = w - textWidth
	case Center:
		x = (w - textWidth) / 2
	}
	height := face.Metrics().Height.Ceil()
	y := height + (h-height)/2

	img, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	text.Draw(img, l.text, face, x, y, l.FontColor)

	for i := 0; i < 9; i++ {
		if i == 4 {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(0, 0, 0, 1)
		dx, dy := i%3-1, i/3-1
		op.GeoM.Translate(float64(dx), float64(dy))
		l.Image.DrawImage(img, op)
	}

	op := &ebiten.DrawImageOptions{}
	l.Image.DrawImage(img, op)
}

// String for fmt.Stringer interface
func (l *Label) String() string {
	return fmt.Sprintf("%d:Label[%s]:%s", l.ID(), l.Rect, string([]rune(l.text)[:4]))
}
