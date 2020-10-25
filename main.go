package main

import (
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
	bg    *ui.Box
	count int
)

// Game is status of game
type Game struct {
	IsDebugPrint bool
}

func init() {
	game = Game{false}
	bg = ui.NewBox(screenWidth, screenHeight, ui.Color("0f0"))
	// menu := menuScreen()
	menu := mouseScreen()
	mainMenu := mainMenu()
	bg.Add(0, 0, mainMenu)
	bg.Add(0, 0, menu)

	ui.SetKeyCallback(ebiten.KeyF4, func(el ui.Element) {
		game.IsDebugPrint = !game.IsDebugPrint
	})
	ui.SetKeyCallback(ebiten.KeyF5, func(el ui.Element) {
		ebiten.SetScreenScale(2)
	})
	ui.SetMouseCallback(ui.LeftClick, func(el ui.Element) {
		log.Println("Foobar!!")
	})
	menu.SetKeyCallback(ebiten.KeyF5, func(el ui.Element) {
		x, y := el.Position()
		el.Move(x+10, y+10)
	})
	menu.SetFocus()
	ui.SetRoot(bg)
}

func (game *Game) Update(screen *ebiten.Image) error {
	ui.Update()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	ui.Draw(screen)
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
