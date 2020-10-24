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
	menu := menuScreen()
	// menu := mouseScreen()
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

func control(screen *ebiten.Image) (err error) {
	ui.Update()
	return
}

func update(screen *ebiten.Image) (err error) {
	return
}

func draw(screen *ebiten.Image) (err error) {
	ui.Draw(screen)
	return
}

func loop(screen *ebiten.Image) (err error) {
	err = control(screen)
	if err != nil {
		return
	}

	err = update(screen)
	if err != nil {
		return
	}

	if ebiten.IsRunningSlowly() {
		return
	}

	err = draw(screen)
	if err != nil {
		return
	}

	if game.IsDebugPrint {
		err = debugPrint(screen)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(loop, screenWidth, screenHeight, 1, "MouseEvent (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
