package gmartini

import (
	"image"
	"image/color"
	_ "image/png"
	"math"

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
			v := minFloat32(max, max*minFloat32((terrain[k]-minZ)/(maxZ-minZ), cutoff)/cutoff)
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
