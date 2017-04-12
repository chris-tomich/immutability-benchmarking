package immutabilitybenchmarking

import "github.com/pkg/errors"

// ImmutableMatrix is a matrix with non-mutating operations.
type ImmutableMatrix struct {
	matrix [][]int
}

// NewImmutableMatrix creates a new matrix with the given initial values.
func NewImmutableMatrix(matrix [][]int) ImmutableMatrix {
	return ImmutableMatrix{matrix: matrix}
}

// EmptyImmutableMatrix createas a new empty matrix with the given dimensions.
func EmptyImmutableMatrix(width int, height int) ImmutableMatrix {
	if width == 0 || height == 0 {
		panic(errors.New("width and height must both be non-zero"))
	}

	m := ImmutableMatrix{
		matrix: make([][]int, height),
	}

	for i := 0; i < len(m.matrix); i++ {
		m.matrix[i] = make([]int, width)
	}

	return m
}

// Equals will compare a matrix against this matrix and return if they are equal.
func (m1 ImmutableMatrix) Equals(m2 ImmutableMatrix) bool {
	if len(m1.matrix) != len(m2.matrix) {
		return false
	}

	if len(m1.matrix[0]) != len(m2.matrix[0]) {
		return false
	}

	for r := 0; r < len(m1.matrix); r++ {
		for c := 0; c < len(m1.matrix[r]); c++ {
			if m1.matrix[r][c] != m2.matrix[r][c] {
				return false
			}
		}
	}

	return true
}

// Add will add the values of a matrix to this matrix.
func (m1 ImmutableMatrix) Add(m2 ImmutableMatrix) (ImmutableMatrix, error) {
	if len(m1.matrix) != len(m2.matrix) {
		return ImmutableMatrix{}, errors.New("width of both matrices are not the same")
	}

	if len(m1.matrix[0]) != len(m2.matrix[0]) {
		return ImmutableMatrix{}, errors.New("height of both matrices are not the same")
	}

	m := EmptyImmutableMatrix(len(m1.matrix), len(m1.matrix[0]))

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] + m2.matrix[r][c]
		}
	}

	return m, nil
}

// Subtract will subtract the values of a matrix from this matrix.
func (m1 ImmutableMatrix) Subtract(m2 ImmutableMatrix) (ImmutableMatrix, error) {
	if len(m1.matrix) != len(m2.matrix) {
		return ImmutableMatrix{}, errors.New("width of both matrices are not the same")
	}

	if len(m1.matrix[0]) != len(m2.matrix[0]) {
		return ImmutableMatrix{}, errors.New("height of both matrices are not the same")
	}

	m := EmptyImmutableMatrix(len(m1.matrix), len(m1.matrix[0]))

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] - m2.matrix[r][c]
		}
	}

	return m, nil
}

// ScalarMultiply will multiply this matrix by a given scalar value.
func (m1 ImmutableMatrix) ScalarMultiply(s int) ImmutableMatrix {
	m := EmptyImmutableMatrix(len(m1.matrix), len(m1.matrix[0]))

	for r := 0; r < len(m1.matrix); r++ {
		for c := 0; c < len(m1.matrix[r]); c++ {
			m.matrix[r][c] = m1.matrix[r][c] * s
		}
	}

	return m
}

// Transpose will transpose this matrix.
func (m1 ImmutableMatrix) Transpose() ImmutableMatrix {
	m := EmptyImmutableMatrix(len(m1.matrix[0]), len(m1.matrix))

	for rt := 0; rt < len(m.matrix); rt++ {
		for ct := 0; ct < len(m1.matrix); ct++ {
			m.matrix[rt][ct] = m1.matrix[ct][rt]
		}
	}

	return m
}

// MatrixMultiply will multiple the given matrix against this matrix.
func (m1 ImmutableMatrix) MatrixMultiply(m2 ImmutableMatrix) (ImmutableMatrix, error) {
	if len(m1.matrix[0]) != len(m2.matrix) {
		return ImmutableMatrix{}, errors.New("the dimensions of the matrices are incompatible, try transposing one first")
	}

	m := EmptyImmutableMatrix(len(m2.matrix[0]), len(m1.matrix))

	for rm := 0; rm < len(m1.matrix); rm++ {
		for cm2 := 0; cm2 < len(m2.matrix[rm]); cm2++ {
			product := 0
			for cm := 0; cm < len(m1.matrix[rm]); cm++ {
				product = product + m1.matrix[rm][cm]*m2.matrix[cm][cm2]
			}
			m.matrix[rm][cm2] = product
		}
	}

	return m, nil
}

