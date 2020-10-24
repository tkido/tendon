package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Frame is nine-patch image
type Frame struct {
	*Image
	Source        image.Image
	Count, Number int
}

// NewFrame make new *ui.Frame
func NewFrame(w, h int, srcImg image.Image, number int) *Frame {
	i := NewImage(w, h, nil)
	n := &Frame{i, srcImg, 0, number}
	n.Self = n
	return n
}

// Reflesh updates internal *ebiten.Image
func (f *Frame) Reflesh() {
	srcImg, _ := ebiten.NewImageFromImage(f.Source, ebiten.FilterDefault)
	srcW, srcH := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()
	frameH := srcH / f.Number
	sub := srcImg.SubImage(image.Rect(0, f.Count*frameH, srcW, (f.Count+1)*frameH))
	f.Image.SetImage(sub)
	f.Image.Reflesh()
}

// SetCount set count
func (f *Frame) SetCount(count int) {
	f.Count = count
	f.dirtySelf()
}

// String for fmt.Stringer interface
func (f *Frame) String() string {
	return fmt.Sprintf("%d:Frame[%s]", f.ID(), f.Box.Rect)
}
