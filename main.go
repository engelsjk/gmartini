package main

import (
	"fmt"
	"image"
	"os"

	"github.com/engelsjk/gmartini/gmartini"
)

func main() {

	var terrainFile string = "data/fuji.png"
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

	mesh := tile.GetMesh(gmartini.OptionMaxError(maxError))

	fmt.Printf("gmartini\n")
	fmt.Printf("***********************\n")
	fmt.Printf("gridsize: %d\n", gridSize)
	fmt.Printf("terrain: %d\n", len(terrain))
	fmt.Printf("max error: %d\n", int(maxError))
	fmt.Printf("***********************\n")
	fmt.Printf("indices: %d\n", len(martini.Indices))
	fmt.Printf("coords: %d\n", len(martini.Coords))
	fmt.Printf("***********************\n")
	fmt.Printf("mesh vertices: %d\n", len(mesh.Vertices))
	fmt.Printf("mesh triangles: %d\n", len(mesh.Triangles))
	fmt.Printf("***********************\n")
}
