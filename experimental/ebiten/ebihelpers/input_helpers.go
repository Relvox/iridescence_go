package ebihelpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func ReadUserInput() string {
	runes := ebiten.AppendInputChars([]rune{})
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		runes = append(runes, '\t')
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
		runes = append(runes, '\n')
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		runes = append(runes, 8)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDelete) {
		runes = append(runes, 127)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		runes = append(runes, 27)
	}

	return string(runes)
}
