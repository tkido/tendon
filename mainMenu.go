package main

import (
	"image/color"

	"github.com/tkido/tendon/ui"
)

func mainMenu() *ui.Box {
	data := []string{"マウス", "キー", "メニュー", "ダイアログ"}
	menu := ui.NewBox(640, 15, ui.Color("ccc"))
	// ui.SetCallback(ebiten.KeyEscape, func() {
	// 	if menu.IsVisible() {
	// 		menu.Hide()
	// 		return
	// 	}
	// 	menu.Show()
	// })
	for i, s := range data {
		label := ui.NewLabel(100, 15, s, FontBold, FontSmall, ui.Center, color.Black, nil)
		label.SetMouseCallback(ui.LeftClick, func(i int) func(el ui.Element) {
			return func(el ui.Element) {
				// result.SetText(data[i])
			}
		}(i))
		label.SetMouseCallback(ui.MouseOver, func(el ui.Element) {
			// op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(20, 0)
			// el.SetDrawImageOptions(op)
		})
		label.SetMouseCallback(ui.MouseOut, func(el ui.Element) {
			// el.SetDrawImageOptions(nil)
		})
		menu.Add(i*120, 0, label)
	}
	return menu
}
