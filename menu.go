package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/tkido/tendon/assets"
	"github.com/tkido/tendon/ui"
	"github.com/fogleman/ease"
	"github.com/hajimehoshi/ebiten"
)

func menuScreen() *ui.Box {
	// images
	f, err := assets.Open("ninepatch.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	data := []string{"あいうえおかきくけこさしすせそたちつてとなにぬねこはひふへほまみむめもやゆよらりるれろわをんabcdeghijklmnopqrstuvwxyz,.!?ABCDEFGHIJKLMNOPQRSTUVWXYZ", "バナナ", "りんご", "オレンジ", "とうもろこし", "焼き肉", "北京ダック", "豚の丸焼き"}
	screen := ui.NewBox(screenWidth, screenHeight, ui.Color("0f0"))
	menu := ui.NewNinepatch(400, 400, png, image.Rect(8, 8, 24, 24))
	screen.Add(10, 10, menu)

	result := ui.NewRichLabel(240, 300, 30, "", FontRegular, FontMedium, ui.Left, ui.Top, color.White, color.Black)
	screen.Add(400, 120, result)

	fadeOut := ui.Animation{
		IsLoop:   false,
		Duration: 30,
		Delay:    0,
		Ease:     ease.InCubic,
		OnAnime: func(el ui.Element, t float64) {
			op := &ebiten.DrawImageOptions{}
			op.ColorM.Scale(1, 1, 1, 1-t)
			el.SetDrawImageOptions(op)
		},
		OnEnd: func(el ui.Element, t float64) {
			el.Hide()
		},
	}

	flashCursor := ui.Animation{
		IsLoop:   true,
		Duration: 30,
		Delay:    0,
		Ease: func(t float64) float64 {
			return math.Sin(2 * math.Pi * t)
		},
		OnAnime: func(el ui.Element, t float64) {
			bgC := color.NRGBA{255, 255, 0, uint8(t*30 + 60)}
			el.SetBackgroundColor(bgC)
		},
		OnEnd: func(el ui.Element, t float64) {
			el.SetBackgroundColor(nil)
		},
	}

	for i, s := range data {
		label := ui.NewLabel(386, 30, s, FontRegular, FontMedium, ui.Left, ui.Color("fff"), color.Transparent)
		label.SetMouseCallback(ui.LeftClick, func(i int) func(el ui.Element) {
			return func(el ui.Element) {
				result.SetText(fmt.Sprintf("%s", data[i]))
			}
		}(i))
		label.SetMouseCallback(ui.MouseOver, func(el ui.Element) {
			el.SetAnimation(flashCursor)
		})
		label.SetMouseCallback(ui.MouseOut, func(el ui.Element) {
			el.StopAnimation()
		})
		menu.Add(8, 8+i*30, label)
		menu.SetMouseCallback(ui.RightClick, func(el ui.Element) {
			el.SetAnimation(fadeOut)
		})
	}
	return screen
}
