package ui

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Box is simple box
type Box struct {
	id                     int
	Rect                   image.Rectangle
	Color                  color.Color
	Canvas                 *ebiten.Image
	isDirty                bool
	Image                  *ebiten.Image
	isDirtySelf            bool
	drawImageOptions       *ebiten.DrawImageOptions
	cachedDrawImageOptions *ebiten.DrawImageOptions
	Parent                 Element
	Children               []Element
	Self                   Element
	visible                bool
	mouseCallbacks
	keyCallbacks
}

// NewBox make new Box
func NewBox(w, h int, c color.Color) *Box {
	b := &Box{
		id:                     gm.nextID(),
		Rect:                   image.Rect(0, 0, w, h),
		Color:                  c,
		Canvas:                 nil,
		isDirty:                true,
		Image:                  nil,
		isDirtySelf:            true,
		drawImageOptions:       &ebiten.DrawImageOptions{},
		cachedDrawImageOptions: &ebiten.DrawImageOptions{},
		Parent:                 nil,
		Children:               []Element{},
		mouseCallbacks:         mouseCallbacks{},
		keyCallbacks:           keyCallbacks{},
		Self:                   nil,
		visible:                true,
	}
	b.Self = b
	return b
}

// ID retruns id
func (b *Box) ID() int {
	return b.id
}

// Show element
func (b *Box) Show() {
	b.visible = true
	if b.Parent != nil {
		b.Parent.dirty()
	}
}

// Hide element
func (b *Box) Hide() {
	b.visible = false
	if b.Parent != nil {
		b.Parent.dirty()
	}
}

// IsVisible return visiblity
func (b *Box) IsVisible() bool {
	return b.visible
}

// Reflesh updates internal *ebiten.Image
func (b *Box) Reflesh() {
	w, h := b.Size()
	b.Image, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	if b.Color == nil || b.Color == color.Transparent {
		return
	}
	b.Image.Fill(b.Color)
}

// dirty set element isDirty
func (b *Box) dirty() {
	if b.isDirty == true {
		return
	}
	b.isDirty = true
	if b.Parent != nil {
		b.Parent.dirty()
	}
}

// dirtySelf set element isDirtySelf
func (b *Box) dirtySelf() {
	b.isDirtySelf = true
	b.dirty()
}

// isDecendantOf prevent stak overflow
func (b *Box) isDecendantOf(child Element) bool {
	if b.ID() == child.ID() {
		return true
	}
	if b.Parent == nil {
		return false
	}
	return b.Parent.isDecendantOf(child)
}

// Add append child element to element
func (b *Box) Add(x, y int, child Element) {
	if b.isDecendantOf(child) {
		log.Fatal("Box.Add: cant't add ancestor as child")
	}
	child.setParent(b.Self)
	b.Children = append(b.Children, child)
	child.Move(x, y)
	b.dirty()
}

// Clear clear children
func (b *Box) Clear() {
	b.Children = []Element{}
	b.dirty()
}

// setParent set parent
func (b *Box) setParent(el Element) {
	b.Parent = el
}

// Move move element. (x, y) is relative position from parent.
func (b *Box) Move(x, y int) {
	w, h := b.Size()
	b.Rect = image.Rect(x, y, x+w, y+h)
	b.dirty()
}

// Position return relative position from parent element
func (b *Box) Position() (x, y int) {
	min := b.Rect.Min
	return min.X, min.Y
}

// Resize resize element
func (b *Box) Resize(w, h int) {
	x, y := b.Position()
	b.Rect = image.Rect(x, y, x+w, y+h)
	b.dirtySelf()
}

// Size get size of element
func (b *Box) Size() (w, h int) {
	s := b.Rect.Size()
	return s.X, s.Y
}

// Draw draw box
func (b *Box) Draw(canvas *ebiten.Image) {
	if !b.visible {
		return
	}
	if b.isDirty {
		b.isDirty = false
		if b.isDirtySelf {
			b.isDirtySelf = false
			b.Self.Reflesh()
		}
		w, h := b.Size()
		b.Canvas, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
		op := &ebiten.DrawImageOptions{}
		b.Canvas.DrawImage(b.Image, op)
		*(b.cachedDrawImageOptions) = *(b.drawImageOptions)
		x, y := b.Position()
		b.cachedDrawImageOptions.GeoM.Translate(float64(x), float64(y))
		for _, c := range b.Children {
			c.Draw(b.Canvas)
		}
	}
	canvas.DrawImage(b.Canvas, b.cachedDrawImageOptions)
}

// SetDrawImageOptions set DrawImageOptions
func (b *Box) SetDrawImageOptions(op *ebiten.DrawImageOptions) {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}
	b.drawImageOptions = op
	b.dirty()
}

// SetAnimation set animation
func (b *Box) SetAnimation(a Animation) {
	gm.SetAnimation(b.Self, a)
}

// StopAnimation stop animation
func (b *Box) StopAnimation() {
	gm.StopAnimation(b.Self)
}

// SetBackgroundColor set background color
func (b *Box) SetBackgroundColor(c color.Color) {
	b.Color = c
	b.dirtySelf()
}

// String for fmt.Stringer
func (b *Box) String() string {
	return fmt.Sprintf("%d:Box[%s]%s", b.ID(), b.Rect, ColorCode(b.Color))
}

// onMouseEvent call callback function if it exists
func (b *Box) onMouseEvent(e MouseEvent) {
	if c, ok := b.mouseCallbacks[e]; ok {
		c(b.Self)
		return
	}
	if LeftClick <= e && e <= MiddleClick && b.Parent != nil {
		b.Parent.onMouseEvent(e)
	}
}

func (b *Box) mouseOn() {
	if gm.OnElement != nil && gm.OnElement != b.Self {
		gm.OnElement.onMouseEvent(MouseOut)
	}
	b.onMouseEvent(MouseOn)
	if gm.OnElement != b.Self {
		b.onMouseEvent(MouseOver)
	}
	gm.OnElement = b.Self
	b.mouseIn()
}

func (b *Box) mouseIn() {
	b.onMouseEvent(MouseIn)
	if _, ok := gm.InElements[b.Self]; !ok {
		b.onMouseEvent(MouseEnter)
	}
	gm.InElements[b.Self] = gm.Now
}

// handleMouseEvent handle mouse event
func (b *Box) handleMouseEvent(ev mouseEvents, origin image.Point, clip image.Rectangle) (handled bool) {
	rect := b.Rect.Add(origin)
	clip = clip.Intersect(rect)
	if !ev.Point.In(clip) {
		return
	}
	// Evaluate children first in reverse order. It is because the child added later is closer to the front.
	for i := len(b.Children) - 1; 0 <= i; i-- {
		if handled := b.Children[i].handleMouseEvent(ev, origin.Add(b.Rect.Min), clip); handled {
			b.mouseIn()
			return true
		}
	}
	// Handle by myself
	b.mouseOn()
	for i := 0; i < 3; i++ {
		down, up, click, doubleClick := LeftDown+MouseEvent(i), LeftUp+MouseEvent(i), LeftClick+MouseEvent(i), LeftDoubleClick+MouseEvent(i)
		switch ev.ButtonEvents[i] {
		case mouseButtonDown:
			b.onMouseEvent(down)
			gm.Downed[i] = &mouseRecord{b, ev.Point, gm.Now}
		case mouseButtonUp:
			b.onMouseEvent(up)
			// isClick?
			if gm.Downed[i] != nil && gm.Downed[i].Element == b && gm.isCloseEnough(ev.Point, gm.Downed[i].Point) {
				if gm.Clicked[i] != nil {
					// isDoubleClick?
					if gm.Clicked[i].Element == b && gm.Now-gm.Clicked[i].Frame <= gm.DoubleClickInterval && gm.isCloseEnough(ev.Point, gm.Clicked[i].Point) {
						b.onMouseEvent(doubleClick)
						gm.Clicked[i] = nil
					} else {
						gm.Clicked[i].Element.onMouseEvent(click)
						gm.Clicked[i] = &mouseRecord{b, ev.Point, gm.Now}
					}
				} else if _, ok := b.mouseCallbacks[doubleClick]; ok {
					gm.Clicked[i] = &mouseRecord{b, ev.Point, gm.Now}
				} else {
					b.onMouseEvent(click)
				}
			}
			gm.Downed[i] = nil
		}
	}
	return true
}
