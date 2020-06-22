package benchmark

import (
	"testing"

	"github.com/engelsjk/gmartini"
)

func BenchmarkMartini(b *testing.B) {

	var gridSize int32 = 513

	for n := 0; n < b.N; n++ {
		gmartini.New(gmartini.OptionGridSize(gridSize))
	}
}
