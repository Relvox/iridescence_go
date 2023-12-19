package ebihelpers

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFontFace(path string, dpi, fontSize float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read font file: %w", err)
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parse open type: %w", err)
	}
	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("face from font: %w", err)

	}
	return font, nil
}

