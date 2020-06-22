package gmartini

import (
	"fmt"
	"image"
	_ "image/png"
	"strings"
)

// DecodeElevation decodes the pixel values of an image.Image (e.g. a Mapbox Terrain RGB raster tile) into a 1D array of heightmap values.
// A backfill option is included to satisfy the martini requirement of a 2^n+1 grid size.
func DecodeElevation(img image.Image, encoding string, addBackfill bool) ([]float32, error) {

	allowedEncodings := make(map[string]bool)
	allowedEncodings["mapbox"] = true
	allowedEncodings["terrarium"] = true

	encodings := []string{}
	for k, v := range allowedEncodings {
		if v {
			encodings = append(encodings, k)
		}
	}

	allowed, ok := allowedEncodings[encoding]
	if !ok {
		return nil, fmt.Errorf("encoding not recognized, must be one of %s", strings.Join(encodings, ", "))
	}
	if !allowed {
		return nil, fmt.Errorf("encoding not allowed, must be one of %s", strings.Join(encodings, ", "))
	}

	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	width := maxX - minX
	height := maxY - minY

	if width != height {
		return nil, fmt.Errorf("width (%d) must equal height (%d)", width, height)
	}

	tileSize := width
	gridSize := width + 1

	terrain := make([]float32, gridSize*gridSize)

	for y := 0; y < tileSize; y++ {
		for x := 0; x < tileSize; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			terrain[y*gridSize+x] = rgbaToTerrain(r, g, b, a, encoding)
		}
	}

	if addBackfill {
		terrain = ComputeBackfill(terrain, gridSize)
	}

	return terrain, nil
}

func ComputeBackfill(arr []float32, gridSize int) []float32 {
	if len(arr) != gridSize*gridSize {
		return arr
	}

	for x := 0; x < gridSize-1; x++ {
		arr[gridSize*(gridSize-1)+x] = arr[gridSize*(gridSize-2)+x] //backfill bottom border
	}
	for y := 0; y < gridSize; y++ {
		arr[gridSize*y+gridSize-1] = arr[gridSize*y+gridSize-2] //backfill right border
	}

	// for i := 0; i < gridSize*gridSize-1; i++ {
	// 	row := i / gridSize
	// 	if i == (row*gridSize + size) { // right border
	// 		arrBackfilled[i] = arr[i-row-1]
	// 	} else if row > (size - 1) { // bottom border
	// 		arrBackfilled[i] = arr[i-row-gridSize]
	// 	} else {
	// 		arrBackfilled[i] = arr[i-row]
	// 	}
	// 	arrBackfilled[gridSize*gridSize-1] = arrBackfilled[gridSize*gridSize-2]
	// }
	return arr
}

// func RescalePosition() {

// func getPixel(img image.Image, x, y int, flipPNG bool) color.Color {
// 	if flipPNG {
// 		return img.At(y, x)
// 	}
// 	return img.At(x, y)
// }

func rgbaToTerrain(r uint32, g uint32, b uint32, a uint32, encoding string) float32 {
	switch encoding {
	case "mapbox":
		return (float32(r>>8)*(256.0*256.0)+float32(g>>8)*256.0+float32(b>>8))/10 - 10000.0
	case "terrarium":
		return (float32(r>>8)*256.0 + float32(g>>8) + float32(b>>8)/256) - 32768.0
	default:
		return 0
	}
}
