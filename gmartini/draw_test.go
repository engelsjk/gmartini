package gmartini

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/fogleman/gg"
)

func TestDrawTerrain(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var imageFile string = "test/terrain.png"

	file, err := os.Open(terrainFile)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Error(err)
	}

	terrain, err := DecodeElevations(img, encoding, true)
	if err != nil {
		t.Error(err)
	}

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTerrain(dc, terrain)
	dc.SavePNG(imageFile)
	t.Logf("test image saved at %s", imageFile)
}

func load(terrainFile, encoding string, gridSize uint, maxError float32) ([]float32, *Mesh) {

	file, err := os.Open(terrainFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	terrain, err := DecodeElevations(img, encoding, true)
	if err != nil {
		panic(err)
	}

	martini, err := New(OptionGridSize(gridSize))
	if err != nil {
		panic(err)
	}

	tile, err := martini.CreateTile(terrain)
	if err != nil {
		panic(err)
	}

	mesh := tile.GetMesh(OptionMaxError(maxError))
	return terrain, mesh
}

func TestDrawVerticesErr5(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 5.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr50(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 50.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr500(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 500.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr5(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 5.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr50(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 50.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr500(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 500.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawAll(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize uint = 513
	var maxError float32 = 50.0
	var imageFile string = "test/martini-%d.png"

	terrain, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTerrain(dc, terrain)
	drawVertices(dc, mesh)
	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}
