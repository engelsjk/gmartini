package main

import (
	"fmt"
	"image"
	"math"
	"os"

	"github.com/engelsjk/gmartini"
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
	fmt.Printf("terrain: %.0f+1 x %.0f+1\n", math.Sqrt(float64(len(terrain)))-1, math.Sqrt(float64(len(terrain)))-1)
	fmt.Printf("max error: %d\n", 30)
	fmt.Printf("mesh vertices: %d\n", mesh.NumVertices)
	fmt.Printf("mesh triangles: %d\n", mesh.NumTriangles)
}
