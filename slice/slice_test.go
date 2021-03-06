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

func BenchmarkMutableMatrix10x10Add(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 10}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrix10x10Add(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 10}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrix10x10Scalar(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 10}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrix10x10Scalar(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 10}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrix10x10Multiply(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 10}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrix10x10Multiply(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 10}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrix10x10Subtract(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 10}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrix10x10Subtract(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 10}
	MatrixSubtractRunner(b, g, 10)
}



func BenchmarkMutableMatrix30x30Add(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 30}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrix30x30Add(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 30}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrix30x30Scalar(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 30}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrix30x30Scalar(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 30}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrix30x30Multiply(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 30}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrix30x30Multiply(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 30}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrix30x30Subtract(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 30}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrix30x30Subtract(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 30}
	MatrixSubtractRunner(b, g, 10)
}



func BenchmarkMutableMatrix90x90Add(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 90}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrix90x90Add(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 90}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrix90x90Scalar(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 90}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrix90x90Scalar(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 90}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrix90x90Multiply(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 90}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrix90x90Multiply(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 90}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrix90x90Subtract(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 90}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrix90x90Subtract(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 90}
	MatrixSubtractRunner(b, g, 10)
}



func BenchmarkMutableMatrix270x270Add(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 270}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrix270x270Add(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 270}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrix270x270Scalar(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 270}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrix270x270Scalar(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 270}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrix270x270Multiply(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 270}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrix270x270Multiply(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 270}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrix270x270Subtract(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 270}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrix270x270Subtract(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 270}
	MatrixSubtractRunner(b, g, 10)
}



func BenchmarkMutableMatrix810x810Add(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 810}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkImmutableMatrix810x810Add(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 810}
	MatrixAddRunner(b, g, 10)
}

func BenchmarkMutableMatrix810x810Scalar(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 810}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkImmutableMatrix810x810Scalar(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 810}
	MatrixScalarRunner(b, g, 10)
}

func BenchmarkMutableMatrix810x810Multiply(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 810}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkImmutableMatrix810x810Multiply(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 810}
	MatrixMultiplyRunner(b, g, 10)
}

func BenchmarkMutableMatrix810x810Subtract(b *testing.B) {
	g := MutableMatrixGenerator{MatrixSize: 810}
	MatrixSubtractRunner(b, g, 10)
}

func BenchmarkImmutableMatrix810x810Subtract(b *testing.B) {
	g := ImmutableMatrixGenerator{MatrixSize: 810}
	MatrixSubtractRunner(b, g, 10)
}
