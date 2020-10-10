package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "math"
	"math/rand"
	"os"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

func main() {

	// constants
	var width, height int = 1000, 1000

	var ALPHA float64 = 2
	var BETA float64 = 1000
	var N int = 2
	rand.Seed(time.Now().UTC().UnixNano())
	var SEED int64 = int64(rand.Float64() * 1000)
	var NOISEFREQ float64 = 950

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	p := perlin.NewPerlin(ALPHA, BETA, N, SEED)

	// white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			alpha := converToAlpha(p.Noise2D(float64(x)/NOISEFREQ, float64(y)/NOISEFREQ))

			var r uint8 = (uint8(x) / uint8((5 * alpha))) * 10
			var g uint8 = (uint8(x) / uint8((5 * alpha))) * 10
			var b uint8 = (uint8(x) / uint8((5 * alpha))) * 10

			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Rendered!")
}

func converToAlpha(val float64) uint8 {
	var OldMin float64 = -1
	var OldMax float64 = 1
	var NewMin float64 = 0
	var NewMax float64 = 255

	var OldRange float64 = (OldMax - OldMin)
	var NewRange float64 = (NewMax - NewMin)
	var NewValue = uint8((((val - OldMin) * NewRange) / OldRange) + NewMin)
	return NewValue
}
