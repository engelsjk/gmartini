package benchmark

import (
	"fmt"
	"image"
	"log"
	"os"
	"testing"
	"time"

	"github.com/engelsjk/gmartini/gmartini"
)

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s: %.03fms", name, float64(elapsed.Nanoseconds())/1000000)
}

func TestExecutionTime(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 30

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

	martini, err := initTileset(gridSize)
	if err != nil {
		panic(err)
	}

	tile, err := createTile(martini, terrain)
	if err != nil {
		panic(err)
	}

	numVertices, numTriangles := generateMesh(tile, maxError, 0, true)
	log.Printf("vertices: %d, triangles: %d\n", numVertices, numTriangles)

	numMeshes := 20
	start := time.Now()
	for i := 0; i < (numMeshes + 1); i++ {
		generateMesh(tile, float32(i), i, false)
	}
	elapsed := time.Since(start)
	log.Printf("%d meshes total: %.03fms", numMeshes, float64(elapsed.Nanoseconds())/1000000)
}

func initTileset(gridSize int32) (*gmartini.Martini, error) {
	defer stopwatch(time.Now(), "init tileset")
	return gmartini.New(gmartini.OptionGridSize(gridSize))
}

func createTile(martini *gmartini.Martini, terrain []float32) (*gmartini.Tile, error) {
	defer stopwatch(time.Now(), "create tile")
	return martini.CreateTile(terrain)
}

func generateMesh(tile *gmartini.Tile, maxError float32, n int, verbose bool) (int32, int) {
	name := fmt.Sprintf("mesh %d", n)
	if verbose {
		name = fmt.Sprintf("mesh (max error = %.0f)", maxError)
	}
	defer stopwatch(time.Now(), name)
	mesh := tile.GetMesh(gmartini.OptionMaxError(maxError))
	return mesh.NumVertices, mesh.NumTriangles
}
