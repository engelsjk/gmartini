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
	log.Printf("%s: %s", name, elapsed)
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

	generateMesh(tile, maxError, 0, true)

	numMeshes := 20
	start := time.Now()
	for i := 0; i < (numMeshes + 1); i++ {
		generateMesh(tile, float32(i), i, false)
	}
	elapsed := time.Since(start)
	log.Printf("%d meshes total: %s", numMeshes, elapsed)
}

func initTileset(gridSize int32) (*gmartini.Martini, error) {
	defer stopwatch(time.Now(), "init tileset")
	return gmartini.New(gmartini.OptionGridSize(gridSize))
}

func createTile(martini *gmartini.Martini, terrain []float32) (*gmartini.Tile, error) {
	defer stopwatch(time.Now(), "create tile")
	return martini.CreateTile(terrain)
}

func generateMesh(tile *gmartini.Tile, maxError float32, n int, verbose bool) {
	var name string
	if verbose {
		name = "mesh"
	} else {
		name = fmt.Sprintf("mesh %d", n)
	}

	defer stopwatch(time.Now(), name)
	mesh := tile.GetMesh(gmartini.OptionMaxError(maxError))

	if verbose {
		name = "mesh"
		log.Printf("vertices: %d, triangles: %d\n", mesh.NumVertices, mesh.NumTriangles)
	}
	return
}
