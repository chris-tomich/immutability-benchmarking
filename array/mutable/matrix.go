package mutable

import (
	"github.com/pkg/errors"
	"github.com/chris-tomich/immutability-benchmarking"
)

// Matrix is a matrix with mutating operations.
type Matrix struct {
	matrix [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int
}

// New creates a new matrix with the given initial values.
func New(matrix [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int) *Matrix {
	return &Matrix{matrix: matrix}
}

// NewEmpty createas a new empty matrix with the given dimensions.
func NewEmpty(width int, height int) (*Matrix, error) {
	if width == 0 || height == 0 {
		return nil, errors.New("width and height must both be non-zero")
	}

	m := &Matrix{}

	return m, nil
}

// Width returns the number of columns in the matrix.
func (m *Matrix) Width() int {
	return len(m.matrix[0])
}

// Height returns the number of rows in the matrix.
func (m *Matrix) Height() int {
	return len(m.matrix)
}

// Get returns the integer at the provided coordinates.
func (m *Matrix) Get(row int, col int) int {
	return m.matrix[row][col]
}

// Equals will compare a matrix against this matrix and return if they are equal.
func (m *Matrix) Equals(m2 immutabilitybenchmarking.Matrix) bool {
	if m.Height() != m2.Height() {
		return false
	}

	if m.Width() != m2.Width() {
		return false
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			if m.matrix[r][c] != m2.Get(r, c) {
				return false
			}
		}
	}

	return true
}

// Add will add the values of a matrix to this matrix.
func (m *Matrix) Add(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m.Height() != m2.Height() {
		return nil, errors.New("width of both matrices are not the same")
	}

	if m.Width() != m2.Width() {
		return nil, errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] + m2.Get(r ,c)
		}
	}

	return m, nil
}

// Subtract will subtract the values of a matrix from this matrix.
func (m *Matrix) Subtract(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m.Height() != m2.Height() {
		return nil, errors.New("width of both matrices are not the same")
	}

	if m.Width() != m2.Width() {
		return nil, errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] - m2.Get(r, c)
		}
	}

	return m, nil
}

// ScalarMultiply will multiply this matrix by a given scalar value.
func (m *Matrix) ScalarMultiply(s int) immutabilitybenchmarking.Matrix {
	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] * s
		}
	}

	return m
}

// Transpose will transpose this matrix.
func (m *Matrix) Transpose() immutabilitybenchmarking.Matrix {
	t := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}

	for rt := 0; rt < len(t); rt++ {
		for ct := 0; ct < len(m.matrix); ct++ {
			t[rt][ct] = m.matrix[ct][rt]
		}
	}

	m.matrix = t

	return m
}

// MatrixMultiply will multiple the given matrix against this matrix.
func (m *Matrix) MatrixMultiply(m2 immutabilitybenchmarking.Matrix) (immutabilitybenchmarking.Matrix, error) {
	if m.Width() != m2.Height() {
		return nil, errors.New("the dimensions of the matrices are incompatible, try transposing one first")
	}

	n := [immutabilitybenchmarking.MatrixHeight][immutabilitybenchmarking.MatrixWidth]int{}

	for rm := 0; rm < m.Height(); rm++ {
		for cm2 := 0; cm2 < m2.Width(); cm2++ {
			product := 0
			for cm := 0; cm < m.Width(); cm++ {
				product = product + m.matrix[rm][cm]*m2.Get(cm, cm2)
			}
			n[rm][cm2] = product
		}
	}

	m.matrix = n

	return m, nil
}
