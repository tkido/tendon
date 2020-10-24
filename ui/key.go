package ui

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

// keyCallbacks is callbacks for key event
type keyCallbacks map[ebiten.Key]Callback

// handleKeyEvent handle key event
func (b *Box) handleKeyEvent(key ebiten.Key) bool {
	if c, ok := b.keyCallbacks[key]; ok {
		c(b.Self)
		return true
	}
	if b.Parent != nil {
		return b.Parent.handleKeyEvent(key)
	}
	return false
}

// SetKeyCallback set callback function for key. set nil means delete.
func (k keyCallbacks) SetKeyCallback(key ebiten.Key, cb Callback) {
	if cb == nil {
		delete(k, key)
		return
	}
	k[key] = cb
	if _, ok := gm.Pressed[key]; !ok {
		gm.Pressed[key] = 0
	}
}

// SetFocus set focus to element
func (b *Box) SetFocus() {
	gm.Focused = b.Self
}

// keyManager manage status of key
type keyManager struct {
	Focused        Element
	Pressed        map[ebiten.Key]uint64
	RepeatInterval int
}

// Update call event
func (m *keyManager) Update() {
	for key, s := range m.Pressed {
		var pressed uint64
		if ebiten.IsKeyPressed(key) {
			pressed = 1
		}
		s = s<<1 | pressed
		m.Pressed[key] = s
		if m.Focused == nil {
			continue
		}
		if s&3 == 1 || s == math.MaxUint64 && gm.Now%m.RepeatInterval == 0 {
			m.Focused.handleKeyEvent(key)
		}
	}
}
