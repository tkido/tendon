package ui

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Update ui
func Update() error {
	gm.Now++
	gm.keyManager.Update()
	// mouse control
	if ev, updated := gm.getMouseEvents(); updated {
		if handled := gm.Root.handleMouseEvent(ev, image.ZP, image.Rect(0, 0, math.MaxInt64, math.MaxInt32)); !handled {
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
	return nil
}

func Draw(screen *ebiten.Image) {
	// update animation
	gm.Animate()
	// start to draw from Root element
	w, h := gm.Root.Size()
	gm.Root.draw(screen, image.Rect(0, 0, w, h))
}

func NewRoot(w, h int, c color.Color) *Box {
	root := NewBox(w, h, c)
	gm.Root = root
	return root
}
