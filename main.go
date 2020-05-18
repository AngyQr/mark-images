package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// Offset is margin top & bottom
const offset = 15

func main() {
	// Open image to marker
	imgResource, _ := os.Open("assets/image.jpg")
	img, _ := jpeg.Decode(imgResource)

	defer imgResource.Close()

	// Get bounds of image
	imgBounds := img.Bounds()
	imgWidth := imgBounds.Max.X

	// Open logo to image
	logoResource, _ := os.Open("assets/logo.png")
	logo, _ := png.Decode(logoResource)

	defer logoResource.Close()

	// Get bounds of logo
	logoBounds := logo.Bounds()
	logoWidth := logoBounds.Max.X

	// Create image output with dimensions of image
	resourceOutput := image.NewRGBA(imgBounds)

	// Coordinates to put logo at image
	point := image.Point{X: imgWidth - (logoWidth + offset), Y: offset}

	// Put logo at image
	draw.Draw(resourceOutput, imgBounds, img, image.ZP, draw.Src)
	draw.Draw(resourceOutput, imgBounds.Add(point), logo, image.ZP, draw.Over)

	// Create image output with logo
	imgOutput, _ := os.Create("output.jpg")
	jpeg.Encode(imgOutput, resourceOutput, &jpeg.Options{Quality: jpeg.DefaultQuality})

	defer imgOutput.Close()
}
