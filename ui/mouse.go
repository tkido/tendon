package ui

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// mouseButtonEvent is move of mouse button
type mouseButtonEvent int

// mouseButtonEvent const
const (
	mouseButtonNone    mouseButtonEvent = 0
	mouseButtonDown                     = 1
	mouseButtonUp                       = 2
	mouseButtonPressed                  = 3
)

// mouseRecord is record of mouse move and event
type mouseRecord struct {
	Element *Box
	Point   image.Point
	Frame   int
}

// mouseEvents is event about mouse action
type mouseEvents struct {
	ButtonEvents [3]mouseButtonEvent
	Point        image.Point
}

// String for fmt.Stringer interface
func (e mouseEvents) String() string {
	return fmt.Sprintf("%v%s", e.ButtonEvents, e.Point)
}

// mouseManager manage status of mouse for ui
type mouseManager struct {
	pressed             [3]byte
	last                mouseEvents
	Downed, Clicked     [3]*mouseRecord
	OnElement           Element
	InElements          map[Element]int
	DoubleClickInterval int // max interval recognized as DoubleClick. Unit is frame(1/60 second)
}

// GetMouseEvent make new mouse event
func (m *mouseManager) getMouseEvents() (es mouseEvents, updated bool) {
	moves := [3]mouseButtonEvent{}
	for i := 0; i < 3; i++ {
		var pressed byte
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton(i)) {
			pressed = 1
		}
		m.pressed[i] = m.pressed[i]<<1 | pressed
		moves[i] = mouseButtonEvent(m.pressed[i] & 3)
	}
	x, y := ebiten.CursorPosition()
	p := image.Point{x, y}
	es = mouseEvents{moves, p}
	if m.last == es {
		return es, false
	}
	m.last = es
	return es, true
}

// isCloseEnough returns whether two points of mouse events are close enough to be regarded as one (single/double) click.
func (m *mouseManager) isCloseEnough(a, b image.Point) bool {
	sub := a.Sub(b)
	x, y := sub.X, sub.Y
	if x*x+y*y <= 4*4 {
		return true
	}
	return false
}
