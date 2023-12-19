package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/relvox/iridescence_go/ebiten/ebihelpers"
)

const (
	WIDTH     = 1920
	HEIGHT    = 1080
	FONT_SIZE = 16
	FONT_PATH = "c:\\Windows\\Fonts\\consolab.ttf"
)

func main() {
	// font, err := ebihelpers.LoadFontFace(FONT_PATH, 10000.0/76.0, float64(FONT_SIZE))
	// if err != nil {
	// 	panic(fmt.Errorf("load game font: %w", err))
	// }

	// ebiten.SetWindowDecorated(false)
	// ebiten.SetWindowSize(g.Size.XY())
	g := ebihelpers.NewTextGridGame(WIDTH, HEIGHT, FONT_PATH, FONT_SIZE, true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
