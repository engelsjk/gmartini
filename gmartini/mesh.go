package gmartini

type Mesh struct {
	MaxError     float32
	NumVertices  int32
	NumTriangles int
	TriIndex     int
	Vertices     []int32
	Triangles    []int32
}

func NewMesh(tile *Tile, opts ...func(*Mesh) error) *Mesh {
	mesh := &Mesh{}
	mesh.NumVertices = 0
	mesh.NumTriangles = 0
	mesh.MaxError = 0

	for _, opt := range opts {
		err := opt(mesh)
		if err != nil {
			return nil
		}
	}

	max := tile.GridSize - 1

	// use an index grid to keep track of vertices that were already used to avoid duplication
	for i := range tile.Indices {
		tile.Indices[i] = 0
	}

	// retrieve mesh in two stages that both traverse the error map:
	// - countElements: find used vertices (and assign each an index), and count triangles (for minimum allocation)
	// - processTriangle: fill the allocated vertices & triangles typed arrays

	mesh.countElements(tile, 0, 0, max, max, max, 0)
	mesh.countElements(tile, max, max, 0, 0, 0, max)

	mesh.Vertices = make([]int32, mesh.NumVertices*2)
	mesh.Triangles = make([]int32, mesh.NumTriangles*3)
	mesh.TriIndex = 0

	mesh.processTriangle(tile, 0, 0, max, max, max, 0)
	mesh.processTriangle(tile, max, max, 0, 0, 0, max)
	return mesh
}

func OptionMaxError(maxError float32) func(*Mesh) error {
	return func(m *Mesh) error {
		m.MaxError = maxError
		return nil
	}
}

func (m *Mesh) countElements(tile *Tile, ax, ay, bx, by, cx, cy int32) {

	mx := (ax + bx) >> 1
	my := (ay + by) >> 1

	middleIndex := my*tile.GridSize + mx

	if absInt32(ax-cx)+absInt32(ay-cy) > 1 && tile.Errors[middleIndex] > m.MaxError {
		m.countElements(tile, cx, cy, ax, ay, mx, my)
		m.countElements(tile, bx, by, cx, cy, mx, my)
	} else {

		aIndex := ay*tile.GridSize + ax
		bIndex := by*tile.GridSize + bx
		cIndex := cy*tile.GridSize + cx

		if tile.Indices[aIndex] == 0 {
			m.NumVertices++
			tile.Indices[aIndex] = m.NumVertices
		}
		if tile.Indices[bIndex] == 0 {
			m.NumVertices++
			tile.Indices[bIndex] = m.NumVertices
		}
		if tile.Indices[cIndex] == 0 {
			m.NumVertices++
			tile.Indices[cIndex] = m.NumVertices
		}
		m.NumTriangles++
	}
}

func (m *Mesh) processTriangle(tile *Tile, ax, ay, bx, by, cx, cy int32) {

	mx := (ax + bx) >> 1
	my := (ay + by) >> 1

	indices := tile.Indices

	middleIndex := my*tile.GridSize + mx

	if absInt32(ax-cx)+absInt32(ay-cy) > 1 && tile.Errors[middleIndex] > m.MaxError {
		// triangle doesn't approximate the surface well enough; drill down further
		m.processTriangle(tile, cx, cy, ax, ay, mx, my)
		m.processTriangle(tile, bx, by, cx, cy, mx, my)
	} else {

		aIndex := ay*tile.GridSize + ax
		bIndex := by*tile.GridSize + bx
		cIndex := cy*tile.GridSize + cx

		// add a triangle
		a := indices[aIndex] - 1
		b := indices[bIndex] - 1
		c := indices[cIndex] - 1

		m.Vertices[2*a] = ax
		m.Vertices[2*a+1] = ay
		m.Vertices[2*b] = bx
		m.Vertices[2*b+1] = by
		m.Vertices[2*c] = cx
		m.Vertices[2*c+1] = cy

		m.Triangles[m.TriIndex] = a
		m.TriIndex++
		m.Triangles[m.TriIndex] = b
		m.TriIndex++
		m.Triangles[m.TriIndex] = c
		m.TriIndex++
	}
}
