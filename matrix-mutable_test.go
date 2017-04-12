package immutabilitybenchmarking

import (
	"testing"
	"math/rand"
)

func TestMutableMatrixMultiplication(t *testing.T) {
	m1 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m2 := NewMutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m1.MatrixMultiply(m2)
	m1EqualsM2 := m1.Equals(NewMutableMatrix(
		[][]int{
			{3, 2340},
			{0, 1000},
		},
	))

	if !m1EqualsM2 {
		t.Fail()
	}

	m3 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
		},
	)

	m4 := NewMutableMatrix(
		[][]int{
			{0, 1000},
			{1, 100},
			{0, 10},
		},
	)

	m3.MatrixMultiply(m4)
	m3EqualsM4 := m3.Equals(NewMutableMatrix(
		[][]int{
			{3, 2340},
		},
	))

	if !m3EqualsM4 {
		t.Fail()
	}

	m5 := NewMutableMatrix(
		[][]int{
			{2, 3, 4},
			{1, 0, 0},
		},
	)

	m6 := NewMutableMatrix(
		[][]int{
			{0},
			{1},
			{0},
		},
	)

	m5.MatrixMultiply(m6)
	m5EqualsM6 := m5.Equals(NewMutableMatrix(
		[][]int{
			{3},
			{0},
		},
	))

	if !m5EqualsM6 {
		t.Fail()
	}
}

func RandomlyGenerateMutableMatrices(size int) (*MutableMatrix, *MutableMatrix) {
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

	return NewMutableMatrix(m1), NewMutableMatrix(m2)
}

func BenchmarkMutableMatrix(b *testing.B) {
	totalMatrices := 1

	mm1 := make([]*MutableMatrix, totalMatrices)
	mm2 := make([]*MutableMatrix, totalMatrices)

	for i := 0; i < totalMatrices; i++ {
		mm1[i], mm2[i] = RandomlyGenerateMutableMatrices(1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < totalMatrices; j++ {
			mm1[j].Add(mm2[j])
			mm1[j].ScalarMultiply(3)
			mm1[j].MatrixMultiply(mm2[j])
			mm1[j].Subtract(mm2[j])
		}
	}
}
