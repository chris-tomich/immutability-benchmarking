package slice

import (
	"testing"
	"github.com/chris-tomich/immutability-benchmarking"
	"github.com/chris-tomich/immutability-benchmarking/slice/mutable"
	"math/rand"
	"github.com/chris-tomich/immutability-benchmarking/slice/immutable"
)

type MatrixGenerator interface {
	Size() int
	GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix)
}

type MutableMatrixGenerator struct {
	MatrixSize int
}

func (m MutableMatrixGenerator) Size() int {
	return m.MatrixSize
}

func (m MutableMatrixGenerator) GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix) {
	m1 := make([][]int, m.Size())
	m2 := make([][]int, m.Size())

	for i := 0; i < m.Size(); i++ {
		m1[i] = make([]int, m.Size())
		m2[i] = make([]int, m.Size())

		for j := 0; j < m.Size(); j++ {
			m1[i][j] = rand.Int()
			m2[i][j] = rand.Int()
		}
	}

	return mutable.New(m1), mutable.New(m2)
}

type ImmutableMatrixGenerator struct {
	MatrixSize int
}

func (m ImmutableMatrixGenerator) Size() int {
	return m.MatrixSize
}

func (m ImmutableMatrixGenerator) GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix) {
	m1 := make([][]int, m.Size())
	m2 := make([][]int, m.Size())

	for i := 0; i < m.Size(); i++ {
		m1[i] = make([]int, m.Size())
		m2[i] = make([]int, m.Size())

		for j := 0; j < m.Size(); j++ {
			m1[i][j] = rand.Int()
			m2[i][j] = rand.Int()
		}
	}

	return immutable.New(m1), immutable.New(m2)
}

func MatrixRunner(b *testing.B, g MatrixGenerator, totalMatrices int) {
	mm1 := make([]immutabilitybenchmarking.Matrix, totalMatrices)
	mm2 := make([]immutabilitybenchmarking.Matrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = g.GenerateMatrix()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j], _ = mm1[j].Add(mm2[j])
			mm1[j] = mm1[j].ScalarMultiply(3)
			mm1[j], _ = mm1[j].MatrixMultiply(mm2[j])
			mm1[j], _ = mm1[j].Subtract(mm2[j])
		}
	}
}

func BenchmarkMutableMatrix(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 50}
	MatrixRunner(b, g, 10)
}

func BenchmarkImmutableMatrix(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 50}
	MatrixRunner(b, g, 10)
}
