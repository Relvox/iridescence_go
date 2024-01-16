package ebihelpers

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen interface {
	SubScreen(x, y, width, height int) Screen
	Fill(clr color.Color)
	DrawImage(image *ebiten.Image)
	DrawRect(x, y, z, w int, clr color.Color)
}

type GoodScreen struct {
	img *ebiten.Image
}

func FromScreen(img *ebiten.Image) GoodScreen {
	return GoodScreen{
		img: img,
	}
}

func (s GoodScreen) SubScreen(x, y, width, height int) Screen {
	y = s.img.Bounds().Dy() - y
	subImage := s.img.SubImage(image.Rect(x, y, x+width, y-height))
	// (subImage.(*ebiten.Image)).Fill(color.RGBA{255, 0, 0, 32})
	return FromScreen(subImage.(*ebiten.Image))
}
func (s GoodScreen) Fill(clr color.Color) {
	s.img.Fill(clr)
}

func (s GoodScreen) DrawImage(image *ebiten.Image) {
	s.img.DrawImage(image, nil)
}

func (s GoodScreen) DrawRect(x, y, z, w int, clr color.Color) {
	width, height := s.img.Bounds().Dx(), s.img.Bounds().Dy()
	pixelCount := 4 * width * height
	pixelData := make([]byte, pixelCount)
	r, g, b, a := clr.RGBA()
	comps := []uint8{uint8(r), uint8(g), uint8(b), uint8(a)}
	s.img.ReadPixels(pixelData)
	var k, l1, l2 int = 0, min(z, w), max(z, w)
	_ = l2
	for ; k < l1; k++ {
		for o, c := range comps {
			pixelData[4*((0)*width+(k))+o] = c
			pixelData[4*((k)*width+(0))+o] = c
			pixelData[4*((w-1)*width+(z-k-1))+o] = c
			pixelData[4*((w-k-1)*width+(z-1))+o] = c
		}
	}
	if l1 == w {
		for ; k < l2; k++ {
			for o, c := range comps {
				pixelData[4*((0)*width+(k))+o] = c
				pixelData[4*((w-1)*width+(z-k-1))+o] = c
			}
		}
	} else {
		for ; k < l2; k++ {
			for o, c := range comps {
				pixelData[4*((w-1-k)*width+(z-1))+o] = c
				pixelData[4*((k)*width+(0))+o] = c
			}
		}
	}
	s.img.WritePixels(pixelData)
}
