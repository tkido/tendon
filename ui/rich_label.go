package ui

import (
	"fmt"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

// VerticalAlign is Vertical Align
type VerticalAlign int

// VerticalAlign
const (
	Top VerticalAlign = iota
	Middle
	Bottom
)

// RichLabel is rich label
type RichLabel struct {
	*Label
	LineHeight int
	VerticalAlign
}

// NewRichLabel makes new *RichLabel
func NewRichLabel(w, h, lineHeight int, text string, fontType FontType, fontSize FontSize, align Align, valign VerticalAlign, color, bgColor color.Color) *RichLabel {
	label := NewLabel(w, h, text, fontType, fontSize, align, color, bgColor)
	l := &RichLabel{label, lineHeight, valign}
	l.Self = l
	return l
}

// Reflesh updates internal *ebiten.Image
func (l *RichLabel) Reflesh() {
	l.Box.Reflesh()
	face := gm.face(l.FontType, l.FontSize)
	w, h := l.Size()

	ss := splitForRichLabel(l.text, face, w)
	total := l.LineHeight * len(ss)
	// log.Printf("%d:%s", len(ss), ss)

	baseY := 0
	switch l.VerticalAlign {
	case Top:
	case Middle:
		baseY = (h - total) / 2
	case Bottom:
		baseY = h - total
	}

	img, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	for i, s := range ss {
		x := 0
		width := font.MeasureString(face, s).Ceil()
		switch l.Align {
		case Left:
		case Right:
			x = w - width
		case Center:
			x = (w - width) / 2
		}
		height := face.Metrics().Height.Ceil()
		y := baseY + (i+1)*l.LineHeight - (l.LineHeight-height)/2
		text.Draw(img, s, face, x, y, l.FontColor)
	}

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
func (l *RichLabel) String() string {
	return fmt.Sprintf("%d:RichLabel[%s]:%s", l.Id(), l.Rect, string([]rune(l.text)[:4]))
}

func splitForRichLabel(s string, f font.Face, w int) []string {
	rs := []rune(s)
	if len(rs) == 0 {
		return []string{}
	}
	rst := []string{}
	cursor := 0
	prevC := rune(-1)
	advance := fixed.Int26_6(0)
	for i, c := range rs {
		// log.Println(string(c))
		// log.Printf("%v", c)
		if prevC >= 0 {
			advance += f.Kern(prevC, c)
		}
		a, ok := f.GlyphAdvance(c)
		if !ok {
			continue
		}
		advance += a
		// log.Printf("i = %d, character = %s, advance = %d", i, string(c), advance.Ceil())
		if advance.Ceil() > w {
			if cursor == i {
				continue
			}
			rst = append(rst, string(rs[cursor:i]))
			cursor = i
			prevC = rune(-1)
			advance, _ = f.GlyphAdvance(c)
		} else if c == '\n' {
			rst = append(rst, string(rs[cursor:i]))
			cursor = i + 1
			prevC = rune(-1)
			advance = fixed.Int26_6(0)
		} else {
			prevC = c
		}
	}
	rst = append(rst, string(rs[cursor:]))
	return rst
}
