package immutable

import (
	"github.com/pkg/errors"
	"github.com/chris-tomich/immutability-benchmarking"
)

// Matrix is an immutable matrix with non-mutating operations.
type Matrix struct {
	matrix [][]int
}

// New creates a new immutable matrix with the given initial values.
func New(matrix [][]int) Matrix {
	return Matrix{matrix: matrix}
}

// NewEmpty createas a new empty matrix with the given dimensions.
func NewEmpty(width int, height int) Matrix {
	if width == 0 || height == 0 {
		panic(errors.New("width and height must both be non-zero"))
	}

	m := Matrix{
		matrix: make([][]int, height),
	}

	for i := 0; i < len(m.matrix); i++ {
		m.matrix[i] = make([]int, width)
	}

	return m
}

// Width returns the number of columns in the matrix.
func (m1 Matrix) Width() int {
	return len(m1.matrix[0])
}

// Height returns the number of rows in the matrix.
func (m1 Matrix) Height() int {
	return len(m1.matrix)
}

// Get returns the integer at the provided coordinates.
func (m1 Matrix) Get(row int, col int) int {
	return m1.matrix[row][col]
}

// Equals will compare a matrix against this matrix and return if they are equal.
func (m1 Matrix) Equals(m2 immutabilitybenchmarking.Matrix) bool {
	if m1.Height() != m2.Height() {
		return false
	}

	if m1.Width() != m2.Width() {
		return false
	}

	for r := 0; r < m1.Height(); r++ {
		for c := 0; c < len(m1.matrix[r]); c++ {
			if m1.matrix[r][c] != m2.Get(r, c) {
				return false
			}
		}
	}

	return true
}

// Add will add the values of a matrix to this matrix.
func (m1 Matrix) Add(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m1.Height() != m2.Height() {
		return Matrix{}, errors.New("width of both matrices are not the same")
	}

	if m1.Width() != m2.Width() {
		return Matrix{}, errors.New("height of both matrices are not the same")
	}

	m := NewEmpty(m1.Height(), m1.Width())

	for r := 0; r < m.Height(); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] + m2.Get(r, c)
		}
	}

	return m, nil
}

// Subtract will subtract the values of a matrix from this matrix.
func (m1 Matrix) Subtract(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m1.Height() != m2.Height() {
		return Matrix{}, errors.New("width of both matrices are not the same")
	}

	if m1.Width() != m2.Width() {
		return Matrix{}, errors.New("height of both matrices are not the same")
	}

	m := NewEmpty(m1.Height(), m1.Width())

	for r := 0; r < m.Height(); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] - m2.Get(r, c)
		}
	}

	return m, nil
}

// ScalarMultiply will multiply this matrix by a given scalar value.
func (m1 Matrix) ScalarMultiply(s int) immutabilitybenchmarking.Matrix {
	m := NewEmpty(m1.Height(), m1.Width())

	for r := 0; r < m1.Height(); r++ {
		for c := 0; c < len(m1.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] * s
		}
	}

	return m
}

// Transpose will transpose this matrix.
func (m1 Matrix) Transpose() immutabilitybenchmarking.Matrix {
	m := NewEmpty(m1.Width(), m1.Height())

	for rt := 0; rt < m.Height(); rt++ {
		for ct := 0; ct < m1.Height(); ct++ {
			m.matrix[rt][ct] = m1.matrix[ct][rt]
		}
	}

	return m
}

// MatrixMultiply will multiple the given matrix against this matrix.
func (m1 Matrix) MatrixMultiply(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m1.Width() != m2.Height() {
		return Matrix{}, errors.New("the dimensions of the matrices are incompatible, try transposing one first")
	}

	m := NewEmpty(m2.Width(), m1.Height())

	for rm := 0; rm < m1.Height(); rm++ {
		for cm2 := 0; cm2 < m2.Width(); cm2++ {
			product := 0
			for cm := 0; cm < len(m1.matrix[rm]); cm++ {
				product = product + m1.matrix[rm][cm]*m2.Get(cm, cm2)
			}
			m.matrix[rm][cm2] = product
		}
	}

	return m, nil
}

