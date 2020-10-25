package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/tendon/ui"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	game  Game
	root  *ui.Box
	count int
)

// Game is status of game
type Game struct {
	IsDebugPrint bool
}

func init() {
	game = Game{false}
	root = ui.NewBox(screenWidth, screenHeight, ui.Color("0f0"))
	menu := menuScreen()
	// menu := mouseScreen()
	mainMenu := mainMenu()
	root.Add(0, 0, mainMenu)
	root.Add(0, 0, menu)

	root.SetKeyCallback(ebiten.KeyF4, func(el ui.Element) {
		game.IsDebugPrint = !game.IsDebugPrint
	})
	root.SetKeyCallback(ebiten.KeyF5, func(el ui.Element) {
		ebiten.SetScreenScale(2)
	})
	root.SetMouseCallback(ui.LeftClick, func(el ui.Element) {
		log.Println("Foobar!!")
	})
	menu.SetKeyCallback(ebiten.KeyF5, func(el ui.Element) {
		x, y := el.Position()
		el.Move(x+10, y+10)
	})
	menu.SetFocus()
}

func (game *Game) Update(screen *ebiten.Image) error {
	ui.Update(root)
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	w, h := root.Size()
	root.Draw(screen, image.Rect(0, 0, w, h))

	if game.IsDebugPrint {
		err := debugPrint(screen)
		if err != nil {
			return
		}
	}
	return
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Charactor Generator")
	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowFloating(true)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
