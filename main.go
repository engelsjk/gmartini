package main

import (
	"fmt"
	"image"
	"os"

	"github.com/engelsjk/gmartini/gmartini"
)

func main() {

	file, err := os.Open("data/fuji.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	terrain, err := gmartini.DecodeElevation(img, "mapbox", true)
	if err != nil {
		panic(err)
	}

	martini, err := gmartini.New(gmartini.OptionGridSize(513))
	if err != nil {
		panic(err)
	}

	tile, err := martini.CreateTile(terrain)
	if err != nil {
		panic(err)
	}

	mesh := tile.GetMesh(gmartini.OptionMaxError(30))

	fmt.Printf("gmartini\n")
	fmt.Printf("***********************\n")
	fmt.Printf("gridsize: %d\n", 513)
	fmt.Printf("terrain: %d\n", len(terrain))
	fmt.Printf("max error: %d\n", 30)
	fmt.Printf("***********************\n")
	fmt.Printf("indices: %d\n", len(martini.Indices))
	fmt.Printf("coords: %d\n", len(martini.Coords))
	fmt.Printf("***********************\n")
	fmt.Printf("mesh vertices: %d\n", len(mesh.Vertices))
	fmt.Printf("mesh triangles: %d\n", len(mesh.Triangles))
	fmt.Printf("***********************\n")
}
