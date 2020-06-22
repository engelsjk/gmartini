package gmartini

type Mesh struct {
	MaxError     float32
	NumVertices  uint32
	NumTriangles uint32
	TriIndex     uint
	Vertices     []uint16
	Triangles    []uint32
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

	size := uint16(tile.GridSize)
	max := size - 1

	// use an index grid to keep track of vertices that were already used to avoid duplication
	for i := range tile.Indices {
		tile.Indices[i] = 0
	}

	// retrieve mesh in two stages that both traverse the error map:
	// - countElements: find used vertices (and assign each an index), and count triangles (for minimum allocation)
	// - processTriangle: fill the allocated vertices & triangles typed arrays

	mesh.countElements(tile, 0, 0, max, max, max, 0)
	mesh.countElements(tile, max, max, 0, 0, 0, max)

	mesh.Vertices = make([]uint16, mesh.NumVertices*2)
	mesh.Triangles = make([]uint32, mesh.NumTriangles*3)
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

func (m *Mesh) countElements(tile *Tile, ax, ay, bx, by, cx, cy uint16) {

	size := uint16(tile.GridSize)

	mx := (ax + bx) >> 1
	my := (ay + by) >> 1

	absAxCx := maxUint16(ax, cx) - minUint16(ax, cx)
	absAyCy := maxUint16(ay, cy) - minUint16(ay, cy)

	middleIndex := uint(my)*uint(size) + uint(mx)

	if absAxCx+absAyCy > 1 && tile.Errors[middleIndex] > m.MaxError {
		m.countElements(tile, cx, cy, ax, ay, mx, my)
		m.countElements(tile, bx, by, cx, cy, mx, my)
	} else {

		aIndex := uint(ay)*uint(size) + uint(ax)
		bIndex := uint(by)*uint(size) + uint(bx)
		cIndex := uint(cy)*uint(size) + uint(cx)

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

func (m *Mesh) processTriangle(tile *Tile, ax, ay, bx, by, cx, cy uint16) {
	mx := (ax + bx) >> 1
	my := (ay + by) >> 1

	indices := tile.Indices
	size := uint16(tile.GridSize)

	absAxCx := maxUint16(ax, cx) - minUint16(ax, cx)
	absAyCy := maxUint16(ay, cy) - minUint16(ay, cy)

	middleIndex := uint(my)*uint(size) + uint(mx)

	if absAxCx+absAyCy > 1 && tile.Errors[middleIndex] > m.MaxError {
		// triangle doesn't approximate the surface well enough; drill down further
		m.processTriangle(tile, cx, cy, ax, ay, mx, my)
		m.processTriangle(tile, bx, by, cx, cy, mx, my)
	} else {

		aIndex := uint(ay)*uint(size) + uint(ax) // x*row+y
		bIndex := uint(by)*uint(size) + uint(bx)
		cIndex := uint(cy)*uint(size) + uint(cx)

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
