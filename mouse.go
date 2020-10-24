package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/tkido/tendon/assets"
	"github.com/tkido/tendon/ui"
)

func onClick(el ui.Element) {
	log.Printf("%s %s", el, "clicked!!")
}
func onDoubleClick(el ui.Element) {
	log.Printf("%s %s", el, "double clicked!!!!")
}
func expand(el ui.Element) {
	w, h := el.Size()
	el.Resize(w+10, h+10)
}

func onMouseOn(el ui.Element) {
	log.Printf("%s %s", el, "MouseOn")
}
func onMouseIn(el ui.Element) {
	log.Printf("%s %s", el, "MouseIn")
}
func onMouseOver(el ui.Element) {
	log.Printf("%s %s", el, "MouseOver")
}
func onMouseOut(el ui.Element) {
	log.Printf("%s %s", el, "MouseOut")
}
func onMouseEnter(el ui.Element) {
	log.Printf("%s %s", el, "MouseEnter")
}
func onMouseLeave(el ui.Element) {
	log.Printf("%s %s", el, "MouseLeave")
}

func mouseScreen() *ui.Box {
	f, err := assets.Open("food_tenpura_ebiten.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	bg := ui.NewBox(screenWidth, screenHeight, ui.Color("0f0"))

	bg.SetMouseCallback(ui.RightClick, onClick)
	bg.SetMouseCallback(ui.RightDoubleClick, onDoubleClick)
	bg.Add(300, 200, ui.NewBox(200, 200, color.Black))

	box1 := ui.NewBox(200, 200, ui.Color("0ff"))
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(2.0, 2.0)
	// op.ColorM.RotateHue(math.Pi)
	// w, h := box1.Size()
	// op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	// op.GeoM.Rotate(math.Pi / 5)
	// op.GeoM.Translate(float64(w)/2, float64(h)/2)
	// box1.SetDrawImageOptions(op)
	for i := -20; i <= 180; i += 100 {
		for j := -20; j <= 180; j += 100 {
			box := ui.NewBox(50, 50, color.NRGBA{0xff, 0x00, 0x00, 0xff})
			box.SetMouseCallback(ui.LeftUp, expand)
			box1.Add(i, j, box)
		}
	}
	bg.Add(200, 100, box1)

	img := ui.NewImage(100, 100, png)
	img.SetMouseCallback(ui.LeftClick, expand)

	box1.Add(-10, 120, img)

	label := ui.NewLabel(screenWidth, 30, "abcdefj", FontRegular, FontMedium, ui.Left, color.White, color.Black)
	label2 := ui.NewLabel(screenWidth, 30, "テストですよ。", FontRegular, FontMedium, ui.Right, color.White, color.Black)
	// label.SetCallback(ui.MouseOn, onMouseOn)
	// label.SetCallback(ui.MouseIn, onMouseIn)
	label.SetMouseCallback(ui.MouseOut, onMouseOut)
	label.SetMouseCallback(ui.MouseOver, onMouseOver)
	label.SetMouseCallback(ui.MouseEnter, onMouseEnter)
	label.SetMouseCallback(ui.MouseLeave, onMouseLeave)
	// bg.SetCallback(ui.MouseOn, onMouseOn)
	// bg.SetCallback(ui.MouseIn, onMouseIn)

	img.SetMouseCallback(ui.RightClick, func(el ui.Element) {
		x, y := el.Position()
		el.Move(x+10, y)
	})

	bg.SetMouseCallback(ui.MouseOut, onMouseOut)
	bg.SetMouseCallback(ui.MouseOver, onMouseOver)
	bg.SetMouseCallback(ui.MouseEnter, onMouseEnter)
	bg.SetMouseCallback(ui.MouseLeave, onMouseLeave)
	bg.Add(10, 10, label)
	bg.Add(10, 40, label2)

	return bg
}
