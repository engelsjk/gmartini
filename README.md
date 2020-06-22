# gmartini

![](test/martini-50.png)

A Go port of the RTIN terrain mesh generator [mapbox/martini](https://github.com/mapbox/martini) by [mourner](https://github.com/mourner).

Additional info on martini can be found at the Observable notebook, "[MARTINI: Real-Time RTIN Terrain Mesh](https://observablehq.com/@mourner/martin-real-time-rtin-terrain-mesh)", written by [mourner](https://github.com/mourner). Original algorithm based on the paper ["Right-Triangulated Irregular Networks" by Will Evans et. al. (1997)](https://www.cs.ubc.ca/~will/papers/rtin.pdf).

Sanity checks of porting correctness and static typing nuance provided by referencing [kylebarron](https://github.com/kylebarron)'s Cython port [pymartini](https://github.com/kylebarron/pymartini).

## Benchmark

A benchmark test is included, showing comparable results to [pymartini](https://github.com/kylebarron/pymartini) in mesh generation but slower in preparation steps.

```bash
go test ./benchmark -run TestExecutionTime -v
```

```
init tileset: 29.507ms
create tile: 9.903ms
mesh (max error = 30): 0.976ms
vertices: 9708, triangles: 19094
mesh 0: 13.710ms
mesh 1: 13.081ms
mesh 2: 11.180ms
mesh 3: 14.529ms
mesh 4: 15.032ms
mesh 5: 12.161ms
mesh 6: 10.418ms
mesh 7: 8.926ms
mesh 8: 7.783ms
mesh 9: 7.231ms
mesh 10: 6.576ms
mesh 11: 5.793ms
mesh 12: 5.020ms
mesh 13: 4.550ms
mesh 14: 4.193ms
mesh 15: 3.833ms
mesh 16: 3.496ms
mesh 17: 3.153ms
mesh 18: 2.945ms
mesh 19: 2.725ms
mesh 20: 2.529ms
20 meshes total: 159.038ms
```
