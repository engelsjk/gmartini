package benchmark

import (
	"image"
	"os"
	"testing"

	"github.com/engelsjk/gmartini"
)

func benchmarkMesh(maxError float32, b *testing.B) {

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

	tile, err := martini.CreateTile(terrain)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		tile.GetMesh(gmartini.OptionMaxError(maxError))
	}
}

func BenchmarkMeshErr5(b *testing.B)   { benchmarkMesh(5, b) }
func BenchmarkMeshErr10(b *testing.B)  { benchmarkMesh(10, b) }
func BenchmarkMeshErr50(b *testing.B)  { benchmarkMesh(50, b) }
func BenchmarkMeshErr100(b *testing.B) { benchmarkMesh(100, b) }
func BenchmarkMeshErr500(b *testing.B) { benchmarkMesh(500, b) }
