package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Element is ebiten UI element
type Element interface {
	Id() int
	isDecendantOf(child Element) bool

	Show()
	Hide()
	IsVisible() bool
	Draw(screen *ebiten.Image)
	Reflesh()
	dirty()
	dirtySelf()

	Move(x, y int)
	Position() (x, y int)
	Resize(w, h int)
	Size() (w, h int)
	SetDrawImageOptions(op *ebiten.DrawImageOptions)
	SetBackgroundColor(c color.Color)
	SetAnimation(anime Animation)
	StopAnimation()

	Add(x, y int, el Element)
	Clear()
	setParent(el Element)

	SetMouseCallback(e MouseEvent, c Callback)
	handleMouseEvent(ev mouseEvents, origin image.Point, clip image.Rectangle) (handled bool)
	onMouseEvent(MouseEvent)

	SetKeyCallback(key ebiten.Key, cb Callback)
	handleKeyEvent(k ebiten.Key) bool
	SetFocus()

	String() string
}

// Texter has internal text as string
type Texter interface {
	SetText(text string)
	Text() (text string)
}

// Imager has internal image source as image.Image
type Imager interface {
	SetImage(srcImage image.Image)
	Source() (srcImage image.Image)
}
