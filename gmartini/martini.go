package gmartini

import (
	"fmt"
)

type Martini struct {
	GridSize           uint
	NumTriangles       int
	NumParentTriangles int
	Indices            []uint32
	Coords             []uint16
}

func New(opts ...func(*Martini) error) (*Martini, error) {
	martini := &Martini{}
	martini.GridSize = 257

	for _, opt := range opts {
		err := opt(martini)
		if err != nil {
			return nil, err
		}
	}

	tileSize := martini.GridSize - 1
	if tileSize&(tileSize-1) == 1 {
		return nil, fmt.Errorf(`expected grid size to be 2^n+1, got %d`, martini.GridSize)
	}

	martini.NumTriangles = int(tileSize*tileSize*2 - 2)
	martini.NumParentTriangles = martini.NumTriangles - int(tileSize*tileSize)

	martini.Indices = make([]uint32, martini.GridSize*martini.GridSize)

	// coordinates for all possible triangles in an RTIN tile
	martini.Coords = make([]uint16, martini.NumTriangles*4)

	// get triangle coordinates from its index in an implicit binary tree
	var id, k uint
	var ax, ay, bx, by, cx, cy uint16

	size := uint16(tileSize)

	for i := 0; i < int(martini.NumTriangles); i++ {
		id = uint(i + 2)

		ax, ay, bx, by, cx, cy = 0, 0, 0, 0, 0, 0
		if id&1 == 1 {
			bx, by, cx = size, size, size // bottom-left triangle
		} else {
			ax, ay, cy = size, size, size // top-right triangle
		}
		for (id >> 1) > 1 {
			id = id >> 1

			mx := (ax + bx) >> 1
			my := (ay + by) >> 1

			if id&1 == 1 { // left half
				bx, by = ax, ay
				ax, ay = cx, cy
			} else { // right half
				ax, ay = bx, by
				bx, by = cx, cy
			}
			cx, cy = mx, my
		}
		k = uint(i * 4)
		martini.Coords[k+0] = ax
		martini.Coords[k+1] = ay
		martini.Coords[k+2] = bx
		martini.Coords[k+3] = by
	}
	return martini, nil
}

func OptionGridSize(gridSize uint) func(*Martini) error {
	return func(m *Martini) error {
		m.GridSize = gridSize
		return nil
	}
}

func (m *Martini) CreateTile(terrain []float32) (*Tile, error) {
	return NewTile(terrain, m)
}
