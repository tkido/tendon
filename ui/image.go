package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Image is simple box
type Image struct {
	*Box
	Source image.Image
}

// NewImage make new *ui.Image
func NewImage(w, h int, srcImg image.Image) *Image {
	b := NewBox(w, h, nil)
	i := &Image{b, srcImg}
	i.Self = i
	return i
}

// Reflesh updates internal *ebiten.Image
func (i *Image) Reflesh() {
	srcImg, _ := ebiten.NewImageFromImage(i.Source, ebiten.FilterDefault)
	w, h := i.Size()
	srcW, srcH := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	i.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	if w != srcW || h != srcH {
		scaleW, scaleH := float64(w)/float64(srcW), float64(h)/float64(srcH)
		op.GeoM.Scale(scaleW, scaleH)
	}
	i.Image.DrawImage(srcImg, op)
}

// SetImage set new source Image
func (i *Image) SetImage(image image.Image) {
	i.Source = image
	i.Dirty()
}

// String for fmt.Stringer interface
func (i *Image) String() string {
	return fmt.Sprintf("%d:Image[%s]", i.Id(), i.Box.Rect)
}
