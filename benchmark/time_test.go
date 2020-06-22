package benchmark

import (
	"fmt"
	"image"
	"log"
	"os"
	"testing"
	"time"

	"github.com/engelsjk/gmartini"
)

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s: %.03fms", name, float64(elapsed.Nanoseconds())/1000000)
}

func TestExecutionTime(t *testing.T) {

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

	martini, err := initTileset(gridSize)
	if err != nil {
		panic(err)
	}

	tile, err := createTile(martini, terrain)
	if err != nil {
		panic(err)
	}

	maxErrors := []float32{0, 2, 5, 10, 20, 30, 50, 75, 100, 250, 500}

	for i := 0; i < len(maxErrors); i++ {
		generateMesh(tile, maxErrors[i])
	}
}

func initTileset(gridSize int32) (*gmartini.Martini, error) {
	defer stopwatch(time.Now(), "init tileset")
	return gmartini.New(gmartini.OptionGridSize(gridSize))
}

func createTile(martini *gmartini.Martini, terrain []float32) (*gmartini.Tile, error) {
	defer stopwatch(time.Now(), "create tile")
	return martini.CreateTile(terrain)
}

func generateMesh(tile *gmartini.Tile, maxError float32) {
	start := time.Now()
	mesh := tile.GetMesh(gmartini.OptionMaxError(maxError))
	elapsed := time.Since(start)

	name := fmt.Sprintf("mesh (max error = %.0f)", maxError)
	vt := fmt.Sprintf("(vertices: %d, triangles: %d)", mesh.NumVertices, mesh.NumTriangles)

	log.Printf("%s: %.03fms %s\n", name, float64(elapsed.Nanoseconds())/1000000, vt)
}
