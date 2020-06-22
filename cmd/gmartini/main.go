package main

import (
	"image"
	"os"

	"github.com/engelsjk/gmartini/gmartini"
)

func main() {
	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50

	file, err := os.Open(terrainFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	terrain, err := gmartini.DecodeElevations(img, encoding, true)
	if err != nil {
		panic(err)
	}

	martini, err := gmartini.New(gmartini.OptionGridSize(gridSize))
	if err != nil {
		panic(err)
	}

	tile, err := martini.CreateTile(terrain)
	if err != nil {
		panic(err)
	}

	_ = tile.GetMesh(gmartini.OptionMaxError(maxError))
}
