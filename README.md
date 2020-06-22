# gmartini

![](test/martini-50.png)

A Go port of the RTIN terrain mesh generator [mapbox/martini](https://github.com/mapbox/martini) by [mourner](https://github.com/mourner).

Additional info on martini can be found at the Observable notebook, "[MARTINI: Real-Time RTIN Terrain Mesh](https://observablehq.com/@mourner/martin-real-time-rtin-terrain-mesh)", written by [mourner](https://github.com/mourner). Original algorithm based on the paper ["Right-Triangulated Irregular Networks" by Will Evans et. al. (1997)](https://www.cs.ubc.ca/~will/papers/rtin.pdf).

Sanity checks of porting correctness and static typing nuance by referencing [kylebarron](https://github.com/kylebarron)'s Cython port [pymartini](https://github.com/kylebarron/pymartini).

## Mesh

A mesh consisting of vertices and triangles can be generated from a terrain PNG image.

```
file, _ := os.Open("data/fuji.png")

img, _, _ := image.Decode(file)

terrain, _ := gmartini.DecodeElevation(img, "mapbox", true)

martini, _ := gmartini.New(gmartini.OptionGridSize(513))

tile, _ := martini.CreateTile(terrain)

mesh := tile.GetMesh(gmartini.OptionMaxError(30))
```
  
## Benchmark

Benchmarking shows comparable results to [pymartini](https://github.com/kylebarron/pymartini) in mesh generation but slower in preparation steps.

```bash
go test ./benchmark -run TestExecutionTime -v
```

```
init tileset:           39.450ms
create tile:            10.201ms
mesh (max error = 0):   14.202ms  (vertices: 261880, triangles: 521768)
mesh (max error = 2):   11.739ms  (vertices: 176260, triangles: 351014)
mesh (max error = 5):   6.633ms   (vertices: 94496,  triangles: 187912)
mesh (max error = 10):  3.313ms   (vertices: 45107,  triangles: 89490)
mesh (max error = 20):  1.329ms   (vertices: 17764,  triangles: 35095)
mesh (max error = 30):  0.779ms   (vertices: 9708,   triangles: 19094)
mesh (max error = 50):  0.361ms   (vertices: 4234,   triangles: 8261)
mesh (max error = 75):  0.191ms   (vertices: 2117,   triangles: 4094)
mesh (max error = 100): 0.120ms   (vertices: 1239,   triangles: 2373)
mesh (max error = 250): 0.032ms   (vertices: 200,    triangles: 359)
mesh (max error = 500): 0.018ms   (vertices: 46,     triangles: 75)
```
