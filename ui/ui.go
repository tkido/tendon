package ui

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// SetRoot set background element
func SetRoot(el Element) {
	gm.Handler.Clear()
	// TODO: clear animation, onItems, defered click, etc.
	gm.Handler.Add(0, 0, el)
	gm.Root = el
}

// Update ui
func Update() {
	gm.Now++
	gm.keyManager.Update()
	// mouse control
	if ev, updated := gm.getMouseEvents(); updated {
		if handled := gm.Handler.handleMouseEvent(ev, image.ZP, image.Rect(0, 0, math.MaxInt64, math.MaxInt32)); !handled {
			if gm.OnElement != nil {
				gm.OnElement.onMouseEvent(MouseOut)
				gm.OnElement = nil
			}
		}
		for k, v := range gm.InElements {
			if v != gm.Now {
				k.onMouseEvent(MouseLeave)
				delete(gm.InElements, k)
			}
		}
	}
	// defered click event callback
	for i := 0; i < 3; i++ {
		if gm.Clicked[i] != nil && gm.Now-gm.Clicked[i].Frame > gm.DoubleClickInterval {
			click := LeftClick + MouseEvent(i)
			gm.Clicked[i].Element.onMouseEvent(click)
			gm.Clicked[i] = nil
		}
	}
}

// Draw ui
func Draw(screen *ebiten.Image) {
	if gm.Root == nil {
		return
	}
	gm.Animate()
	w, h := gm.Root.Size()
	rect := image.Rect(0, 0, w, h)
	gm.Root.Draw(screen, rect)
}

// SetKeyCallback set callback function for key. set nil means delete.
func SetKeyCallback(key ebiten.Key, cb Callback) {
	gm.Handler.SetKeyCallback(key, cb)
}

// SetMouseCallback set callback function for mouse. set nil means delete.
func SetMouseCallback(ev MouseEvent, cb Callback) {
	gm.Handler.SetMouseCallback(ev, cb)
}

// ClearFocus clear focus
func ClearFocus() {
	gm.Handler.SetFocus()
}
