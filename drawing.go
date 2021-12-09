package BioHits

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/psykhi/wordclouds"
)

// drawWordCloud function draws a word cloud image for a word frequency map
func DrawWordCloud(wordFreq map[string]int, filename, fontFile string) {
	colorsRGBA := []color.RGBA{
		{136, 202, 94, 255},
		{66, 151, 160, 255},
		{229, 127, 132, 255},
		{47, 80, 97, 255},
	}
	// set colors
	colors := make([]color.Color, 0)
	for _, c := range colorsRGBA {
		colors = append(colors, c)
	}
	w := wordclouds.NewWordcloud(
		wordFreq,
		wordclouds.FontFile(fontFile),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.Colors(colors),
	)

	img := w.Draw()

	f, err := os.Create(filename)
	if err != nil {
		panic("Cannot create a png file for word cloud.")
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	fmt.Print("Successfully generated WordCloud image!")
	if err != nil {
		panic("Cannot encode image to png.")
	}
}
