package ui

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

// gm is the singleton instance of general manager
var gm = &generalManager{
	Count:   -1,
	Now:     0,
	Root:    nil,
	Handler: nil,
	mouseManager: mouseManager{
		Downed:              [3]*mouseRecord{},
		Clicked:             [3]*mouseRecord{},
		InElements:          map[Element]int{},
		DoubleClickInterval: 15,
	},
	keyManager: keyManager{
		Focused:        nil,
		Pressed:        map[ebiten.Key]uint64{},
		RepeatInterval: 6,
	},
	fontManager: fontManager{
		map[FontType]fontData{},
	},
	animationManager: animationManager{
		map[int]*animation{},
	},
}

func init() {
	gm.Handler = NewBox(math.MaxInt64, math.MaxInt64, nil)
	ClearFocus()
}

// generalManager is manager of internal status of ui
type generalManager struct {
	Count   int
	Now     int
	Root    Element
	Handler Element // global key and mouse handler. implementation is just an empty Box that is always set as the parent of Root.
	animationManager
	fontManager
	keyManager
	mouseManager
}

// nextID retruns next element id
func (m *generalManager) nextID() int {
	m.Count++
	return m.Count
}
