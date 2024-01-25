package ebihelpers

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Screen interface {
	Bounds() image.Rectangle

	SubScreen(rect image.Rectangle) Screen
	Fill(clr color.Color)
	FillRect(rect image.Rectangle, clr color.Color)
	DrawImage(image *ebiten.Image)
	DrawOutline(rect image.Rectangle, clr color.Color)
}

type GoodScreen struct {
	img *ebiten.Image
}

func FromScreen(img *ebiten.Image) GoodScreen { return GoodScreen{img: img} }

func (s GoodScreen) Bounds() image.Rectangle { return s.img.Bounds() }

func (s GoodScreen) Fill(clr color.Color) { s.img.Fill(clr) }

func (s GoodScreen) SubScreen(rect image.Rectangle) Screen {
	rect = image.Rect(
		rect.Min.X, s.Bounds().Max.Y-rect.Max.Y,
		rect.Max.X, s.Bounds().Max.Y-rect.Min.Y)
	return FromScreen(s.img.SubImage(rect).(*ebiten.Image))
}

func (s GoodScreen) FillRect(rect image.Rectangle, clr color.Color) {
	var x int = s.img.Bounds().Min.X + rect.Min.X
	var y int = (s.img.Bounds().Max.Y - 1) - rect.Min.Y
	var width, height int = rect.Dx(), rect.Dy()
	vector.DrawFilledRect(s.img, float32(x), float32(y-height), float32(width), float32(height), clr, false)
}

func (s GoodScreen) DrawImage(image *ebiten.Image) {
	geom := ebiten.GeoM{}
	rect := s.Bounds()
	geom.Scale(float64(rect.Dx())/float64(image.Bounds().Dx()), float64(rect.Dy())/float64(image.Bounds().Dy()))
	geom.Translate(float64(rect.Min.X), float64(rect.Min.Y))
	s.img.DrawImage(image, &ebiten.DrawImageOptions{GeoM: geom})
}

func (s GoodScreen) DrawOutline(rect image.Rectangle, clr color.Color) {
	x0, y0 := rect.Min.X, s.Bounds().Max.Y-rect.Min.Y
	x1, y1 := rect.Max.X, s.Bounds().Max.Y-rect.Max.Y
	vector.StrokeLine(s.img, float32(x0), float32(y0), float32(x1), float32(y0), 1, clr, false)
	vector.StrokeLine(s.img, float32(x0+1), float32(y0), float32(x0+1), float32(y1), 1, clr, false)
	vector.StrokeLine(s.img, float32(x1), float32(y1+1), float32(x0), float32(y1+1), 1, clr, false)
	vector.StrokeLine(s.img, float32(x1), float32(y1), float32(x1), float32(y0), 1, clr, false)
}
