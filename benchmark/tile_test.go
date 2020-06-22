package benchmark

import (
	"image"
	"os"
	"testing"

	"github.com/engelsjk/gmartini"
)

func BenchmarkTile(b *testing.B) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513

	file, err := os.Open(terrainFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	terrain, err := gmartini.DecodeElevation(img, encoding, true)
	if err != nil {
		panic(err)
	}

	martini, err := gmartini.New(gmartini.OptionGridSize(gridSize))
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		martini.CreateTile(terrain)
	}
}
