package BioHits

import (
	"image/color"
	"image/png"
	"os"

	"github.com/psykhi/wordclouds"
)

// drawWordCloud function draws a word cloud image for a word frequency map
func drawWordCloud(wordFreq map[string]int, filename, fontFile string) {
	colorsRGBA := []color.RGBA{
		{0, 109, 91, 0xff},
		{255, 255, 240, 0xff},
		{255, 127, 80, 0xff},
		{160, 175, 183, 0xff},
	}
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
	//fmt.Println("here")
	if err != nil {
		panic("Cannot encode image to png.")
	}
}
