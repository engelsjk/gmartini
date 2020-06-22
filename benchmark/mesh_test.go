package benchmark

import (
	"image"
	"os"
	"testing"

	"github.com/engelsjk/gmartini/gmartini"
)

func benchmarkMesh(maxError float32, b *testing.B) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513

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

	for n := 0; n < b.N; n++ {
		tile.GetMesh(gmartini.OptionMaxError(maxError))
	}

	// fmt.Printf("gmartini\n")
	// fmt.Printf("***********************\n")
	// fmt.Printf("gridsize: %d\n", gridSize)
	// fmt.Printf("terrain: %d\n", len(terrain))
	// fmt.Printf("max error: %d\n", int(maxError))
	// fmt.Printf("***********************\n")
	// fmt.Printf("indices: %d\n", len(martini.Indices))
	// fmt.Printf("coords: %d\n", len(martini.Coords))
	// fmt.Printf("***********************\n")
	// fmt.Printf("mesh vertices: %d\n", len(mesh.Vertices))
	// fmt.Printf("mesh triangles: %d\n", len(mesh.Triangles))
	// fmt.Printf("***********************\n")

}

func BenchmarkMeshErr5(b *testing.B)   { benchmarkMesh(5, b) }
func BenchmarkMeshErr10(b *testing.B)  { benchmarkMesh(10, b) }
func BenchmarkMeshErr50(b *testing.B)  { benchmarkMesh(50, b) }
func BenchmarkMeshErr100(b *testing.B) { benchmarkMesh(100, b) }
func BenchmarkMeshErr500(b *testing.B) { benchmarkMesh(500, b) }
