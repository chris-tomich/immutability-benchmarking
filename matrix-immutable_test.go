package immutabilitybenchmarking

import (
	"testing"
	"math/rand"
)

func TestImmutableMatrixMultiplication(t *testing.T) {
	m1 := NewImmutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m2 := NewImmutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m3, _ := m1.MatrixMultiply(m2)
	isCorrect := m3.Equals(NewImmutableMatrix(
		[][]int{
			{3, 2340},
			{0, 1000},
		},
	))

	if !isCorrect {
		t.Fail()
	}

	m4 := NewImmutableMatrix(
		[][]int{
			{2, 3, 4},
		},
	)

	m5 := NewImmutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m6, _ := m4.MatrixMultiply(m5)
	isCorrect = m6.Equals(NewImmutableMatrix(
		[][]int{
			{3, 2340},
		},
	))

	if !isCorrect {
		t.Fail()
	}

	m7 := NewImmutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m8 := NewImmutableMatrix(
		[][]int{
			{0},
			{1},
			{0},
		},
	)

	m9, _ := m7.MatrixMultiply(m8)
	isCorrect = m9.Equals(NewImmutableMatrix(
		[][]int{
			{3},
			{0},
		},
	))

	if !isCorrect {
		t.Fail()
	}
}

func RandomlyGenerateImmutableMatrices(size int) (ImmutableMatrix, ImmutableMatrix) {
	m1 := make([][]int, size)
	m2 := make([][]int, size)

	for i := 0; i < size; i++ {
		m1[i] = make([]int, size)
		m2[i] = make([]int, size)

		for j := 0; j < size; j++ {
			m1[i][j] = rand.Int()
			m2[i][j] = rand.Int()
		}
	}

	return NewImmutableMatrix(m1), NewImmutableMatrix(m2)
}

func BenchmarkImmutableMatrix(b *testing.B) {
	totalMatrices := 1

	mm1 := make([]ImmutableMatrix, totalMatrices)
	mm2 := make([]ImmutableMatrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = RandomlyGenerateImmutableMatrices(1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mresult, _ := mm1[j].Add(mm2[j])
			mresult = mresult.ScalarMultiply(3)
			mresult, _ = mresult.MatrixMultiply(mm2[j])
			mresult, _ = mresult.Subtract(mm2[j])
		}
	}
}

