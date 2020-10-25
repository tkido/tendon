package ui

import (
	"image"
	"math"
)

// Update ui
func Update(root Element) {
	gm.Now++
	gm.keyManager.Update()
	// mouse control
	if ev, updated := gm.getMouseEvents(); updated {
		if handled := root.handleMouseEvent(ev, image.ZP, image.Rect(0, 0, math.MaxInt64, math.MaxInt32)); !handled {
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
	// update animation
	gm.Animate()
}
