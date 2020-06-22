package gmartini

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"os"
	"testing"

	"github.com/engelsjk/cturbo"
	"github.com/fogleman/gg"
)

func drawTerrain(dc *gg.Context, terrain []float32) {
	var cutoff, max float32 = 1.0, 1.0
	size := int(math.Sqrt(float64(len(terrain))))

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	maxZ, minZ := maxminFloat32(terrain)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			k := y*size + x
			v := minFloat32v(max, max*minFloat32v((terrain[k]-minZ)/(maxZ-minZ), cutoff)/cutoff)
			r, g, b, a := cturbo.Map(float64(v), 255)
			img.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	dc.DrawImage(img, 0, 0)
}

func drawVertices(dc *gg.Context, mesh *Mesh) {
	dc.SetRGB(0, 0, 0)
	for i := 0; i < (len(mesh.Vertices) - 2); i += 2 {
		dc.DrawCircle(float64(mesh.Vertices[i]), float64(mesh.Vertices[i+1]), 0.5)
		dc.Fill()
	}
}

func drawTriangles(dc *gg.Context, mesh *Mesh) {
	dc.ClearPath()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(0.5)
	for i := 0; i < (len(mesh.Triangles) - 3); i += 3 {
		a, b, c := mesh.Triangles[i], mesh.Triangles[i+1], mesh.Triangles[i+2]
		ax, ay := float64(mesh.Vertices[2*a]), float64(mesh.Vertices[2*a+1])
		bx, by := float64(mesh.Vertices[2*b]), float64(mesh.Vertices[2*b+1])
		cx, cy := float64(mesh.Vertices[2*c]), float64(mesh.Vertices[2*c+1])
		dc.MoveTo(ax, ay)
		dc.LineTo(bx, by)
		dc.LineTo(cx, cy)
		dc.LineTo(ax, ay)
	}
	dc.Stroke()
}

func TestDrawTerrain(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var imageFile string = "../test/terrain.png"

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

func load(terrainFile, encoding string, gridSize int32, maxError float32) ([]float32, *Mesh) {

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

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 5.0
	var imageFile string = "../test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr50(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50.0
	var imageFile string = "../test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr500(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 500.0
	var imageFile string = "../test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr5(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
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

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
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

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 500.0
	var imageFile string = "../test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawAll(t *testing.T) {

	var terrainFile string = "../data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50.0
	var imageFile string = "../test/martini-%d.png"

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
