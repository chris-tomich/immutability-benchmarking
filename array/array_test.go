package array

import (
	"testing"
	"github.com/chris-tomich/immutability-benchmarking"
	"github.com/chris-tomich/immutability-benchmarking/array/mutable"
	"math/rand"
	"github.com/chris-tomich/immutability-benchmarking/array/immutable"
)

type MatrixGenerator interface {
	GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix)
}

type MutableMatrixGenerator struct {}

func (MutableMatrixGenerator) GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix) {
	m1 := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}
	m2 := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}

	for i := 0; i < immutabilitybenchmarking.MatrixHeight; i++ {
		for j := 0; j < immutabilitybenchmarking.MatrixWidth; j++ {
			m1[i][j] = rand.Int()
			m2[i][j] = rand.Int()
		}
	}

	return mutable.New(m1), mutable.New(m2)
}

type ImmutableMatrixGenerator struct {}

func (ImmutableMatrixGenerator) GenerateMatrix() (immutabilitybenchmarking.Matrix, immutabilitybenchmarking.Matrix) {
	m1 := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}
	m2 := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}

	for i := 0; i < immutabilitybenchmarking.MatrixHeight; i++ {
		for j := 0; j < immutabilitybenchmarking.MatrixWidth; j++ {
			m1[i][j] = rand.Int()
			m2[i][j] = rand.Int()
		}
	}

	return immutable.New(m1), immutable.New(m2)
}

func MatrixAddRunner(b *testing.B, g MatrixGenerator, totalMatrices int) {
	mm1 := make([]immutabilitybenchmarking.Matrix, totalMatrices)
	mm2 := make([]immutabilitybenchmarking.Matrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = g.GenerateMatrix()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j], _ = mm1[j].Add(mm2[j])
		}
	}
}

func MatrixScalarRunner(b *testing.B, g MatrixGenerator, totalMatrices int) {
	mm1 := make([]immutabilitybenchmarking.Matrix, totalMatrices)
	mm2 := make([]immutabilitybenchmarking.Matrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = g.GenerateMatrix()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j] = mm1[j].ScalarMultiply(3)
		}
	}
}

func MatrixMultiplyRunner(b *testing.B, g MatrixGenerator, totalMatrices int) {
	mm1 := make([]immutabilitybenchmarking.Matrix, totalMatrices)
	mm2 := make([]immutabilitybenchmarking.Matrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = g.GenerateMatrix()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j], _ = mm1[j].MatrixMultiply(mm2[j])
		}
	}
}

func MatrixSubtractRunner(b *testing.B, g MatrixGenerator, totalMatrices int) {
	mm1 := make([]immutabilitybenchmarking.Matrix, totalMatrices)
	mm2 := make([]immutabilitybenchmarking.Matrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = g.GenerateMatrix()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j], _ = mm1[j].Subtract(mm2[j])
		}
	}
}

func BenchmarkMutableMatrixAdd(b *testing.B) {
	g := MutableMatrixGenerator{}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrixAdd(b *testing.B) {
	g := ImmutableMatrixGenerator{}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrixScalar(b *testing.B) {
	g := MutableMatrixGenerator{}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrixScalar(b *testing.B) {
	g := ImmutableMatrixGenerator{}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrixMultiply(b *testing.B) {
	g := MutableMatrixGenerator{}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrixMultiply(b *testing.B) {
	g := ImmutableMatrixGenerator{}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrixSubtract(b *testing.B) {
	g := MutableMatrixGenerator{}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrixSubtract(b *testing.B) {
	g := ImmutableMatrixGenerator{}
	MatrixSubtractRunner(b, g, 10)
}
