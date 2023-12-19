package ebihelpers

import (
	"fmt"
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/relvox/iridescence_go/geom"
	iritext "github.com/relvox/iridescence_go/text"
)

type TextOrientation uint8

const (
	OriginTop     TextOrientation = 0
	OriginBottom  TextOrientation = 1
	RunFromOrigin TextOrientation = 0
	StayAtOrigin  TextOrientation = 2
)

const (
	ZOOM_STEP = 2
)

type TextGridGame struct {
	fontPath        string
	allowInput      bool
	Size            geom.Point2
	FontFace        font.Face
	FontSize        int
	Text            *iritext.TextGrid
	TextOrientation TextOrientation
}

func NewTextGridGame(width, height int, fontPath string, fontSize int, allowInput bool) TextGridGame {
	font, err := LoadFontFace(fontPath, 10000.0/76.0, float64(fontSize))
	if err != nil {
		panic(fmt.Errorf("load game font: %w", err))
	}
	return TextGridGame{
		fontPath:   fontPath,
		allowInput: allowInput,
		Size:       [2]int{width, height},
		FontFace:   font,
		FontSize:   fontSize,
		Text:       iritext.NewTextGrid(),
	}
}

func (g *TextGridGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if outsideWidth != g.Size.X() || outsideHeight != g.Size.Y() {
		g.Size = geom.Point2{outsideWidth, outsideHeight}
	}
	x, y := g.Size.XY()
	return x, y
}

func (g *TextGridGame) Update() error {
	eikjp := inpututil.IsKeyJustPressed

	if eikjp(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		switch {
		case eikjp(ebiten.KeyNumpadAdd) || eikjp(ebiten.KeyEqual):
			g.FontSize += ZOOM_STEP
			if font, err := LoadFontFace(g.fontPath, 10000.0/76.0, float64(g.FontSize)); err == nil {
				g.FontFace = font
			}
		case eikjp(ebiten.KeyNumpadSubtract) || eikjp(ebiten.KeyMinus):
			g.FontSize -= ZOOM_STEP
			if font, err := LoadFontFace(g.fontPath, 10000.0/76.0, float64(g.FontSize)); err == nil {
				g.FontFace = font
			}
		case eikjp(ebiten.Key0) || eikjp(ebiten.KeyNumpad0):
			g.FontSize = 0
			if font, err := LoadFontFace(g.fontPath, 10000.0/76.0, float64(g.FontSize)); err == nil {
				g.FontFace = font
			}
		}
	}
	if g.allowInput {
		input := ReadUserInput()
		g.Text.AppendString(input)
	}
	return nil
}

func (g *TextGridGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.Transparent)
	fontSize := g.FontSize
	w, h := g.Size.XY()
	// for x := 0; x <= w-2*fontSize; x++ {
	// 	for y := 0; y <= h-2*fontSize; y++ {
	// 		if x%(fontSize*2) == 0 || y%(fontSize*2) == 0 {
	// 			vector.DrawFilledRect(screen, float32(fontSize+x), float32(fontSize+y), 1, 1, color.RGBA{G: 255, A: 128}, false)
	// 		} else if x%fontSize == 0 || y%fontSize == 0 {
	// 			vector.DrawFilledRect(screen, float32(fontSize+x), float32(fontSize+y), 1, 1, color.RGBA{G: 128, A: 64}, false)
	// 		}
	// 	}
	// }
	origin, flip := g.TextOrientation&OriginBottom == OriginBottom, g.TextOrientation&StayAtOrigin == StayAtOrigin
	boxedLines := g.Text.BoxedLines(w, h, fontSize, 10, origin != flip)
	if flip {
		slices.Reverse(boxedLines)
	}
	for i, line := range boxedLines {
		y := fontSize * 2 * (i + 1)
		if origin {
			y = h - y
		}
		text.Draw(screen, line, g.FontFace, fontSize, y, color.RGBA{R: 240, G: 240, B: 240, A: 255})
	}
}
