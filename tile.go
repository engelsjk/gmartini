package gmartini

import (
	"fmt"
)

type Tile struct {
	GridSize           int32
	NumTriangles       int
	NumParentTriangles int
	Indices            []int32
	Coords             []int32
	Terrain            []float32
	Errors             []float32
}

// NewTile instantiates a new Tile instance of the generated RTIN hierarchy using a specified terrain map and an initialized Martini instance.
// This hierarchy is an array of error values determined by interpolated height values relative to the specified terrain map.
func NewTile(terrain []float32, martini *Martini) (*Tile, error) {
	tile := &Tile{}

	size := martini.GridSize

	if len(terrain) != int(size*size) {
		return nil, fmt.Errorf(`expected terrain data of length %d (%d x %d), got %d`, size*size, size, size, len(terrain))
	}

	tile.Terrain = terrain
	tile.Errors = make([]float32, len(terrain))

	tile.GridSize = martini.GridSize
	tile.NumTriangles = martini.NumTriangles
	tile.NumParentTriangles = martini.NumParentTriangles

	tile.Indices = martini.Indices
	tile.Coords = martini.Coords

	tile.update()
	return tile, nil
}

func (t *Tile) update() {

	var k int
	var ax, ay, bx, by, cx, cy, mx, my int32
	var interpolatedHeight, middleError float32
	var middleIndex, leftChildIndex, rightChildIndex int32
	var aIndex, bIndex int32

	for i := t.NumTriangles - 1; i >= 0; i-- {
		k = i * 4
		ax = t.Coords[k+0]
		ay = t.Coords[k+1]
		bx = t.Coords[k+2]
		by = t.Coords[k+3]
		mx = (ax + bx) >> 1
		my = (ay + by) >> 1
		cx = mx + my - ay
		cy = my + ax - mx

		// calculate error in the middle of the long edge of the triangle

		aIndex = ay*t.GridSize + ax
		bIndex = by*t.GridSize + bx

		interpolatedHeight = (t.Terrain[aIndex] + t.Terrain[bIndex]) / 2
		middleIndex = my*t.GridSize + mx
		middleError = absFloat32(interpolatedHeight - t.Terrain[middleIndex])

		t.Errors[middleIndex] = maxFloat32x2(t.Errors[middleIndex], middleError)

		if i < t.NumParentTriangles { // bigger triangles; accumulate error with children
			leftChildIndex = ((ay+cy)>>1)*t.GridSize + ((ax + cx) >> 1)
			rightChildIndex = ((by+cy)>>1)*t.GridSize + ((bx + cx) >> 1)
			t.Errors[middleIndex] = maxFloat32x3(t.Errors[middleIndex], t.Errors[leftChildIndex], t.Errors[rightChildIndex])
		}
	}
}

// GetMesh generates a new mesh of vertices and triangles for the Tile with a max error (default 0).
func (t *Tile) GetMesh(opts ...func(*Mesh) error) *Mesh {
	return NewMesh(t, opts...)
}
