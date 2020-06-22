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

func NewTile(terrain []float32, martini *Martini) (*Tile, error) {
	tile := &Tile{}

	size := int32(martini.GridSize)

	// todo: allow ndarray as input

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

	tile.Update()
	return tile, nil
}

func (t *Tile) Update() {

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

		t.Errors[middleIndex] = maxFloat32(t.Errors[middleIndex], middleError)

		if i < t.NumParentTriangles { // bigger triangles; accumulate error with children
			leftChildIndex = ((ay+cy)>>1)*t.GridSize + ((ax + cx) >> 1)
			rightChildIndex = ((by+cy)>>1)*t.GridSize + ((bx + cx) >> 1)
			t.Errors[middleIndex] = maxFloat32(t.Errors[middleIndex], t.Errors[leftChildIndex], t.Errors[rightChildIndex])
		}
	}
}

func (t *Tile) GetMesh(opts ...func(*Mesh) error) *Mesh {
	return NewMesh(t, opts...)
}
