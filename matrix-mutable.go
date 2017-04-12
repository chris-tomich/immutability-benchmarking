package immutabilitybenchmarking

import "github.com/pkg/errors"

// MutableMatrix is a matrix with mutating operations.
type MutableMatrix struct {
	matrix [][]int
}

// NewMutableMatrix creates a new matrix with the given initial values.
func NewMutableMatrix(matrix [][]int) *MutableMatrix {
	return &MutableMatrix{matrix: matrix}
}

// EmptyMutableMatrix createas a new empty matrix with the given dimensions.
func EmptyMutableMatrix(width int, height int) (*MutableMatrix, error) {
	if width == 0 || height == 0 {
		return nil, errors.New("width and height must both be non-zero")
	}

	m := &MutableMatrix{
		matrix: make([][]int, height),
	}

	for i := 0; i < len(m.matrix); i++ {
		m.matrix[i] = make([]int, width)
	}

	return m, nil
}

// Equals will compare a matrix against this matrix and return if they are equal.
func (m *MutableMatrix) Equals(m2 *MutableMatrix) bool {
	if len(m.matrix) != len(m2.matrix) {
		return false
	}

	if len(m.matrix[0]) != len(m2.matrix[0]) {
		return false
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			if m.matrix[r][c] != m2.matrix[r][c] {
				return false
			}
		}
	}

	return true
}

// Add will add the values of a matrix to this matrix.
func (m *MutableMatrix) Add(m2 *MutableMatrix) error {
	if len(m.matrix) != len(m2.matrix) {
		return errors.New("width of both matrices are not the same")
	}

	if len(m.matrix[0]) != len(m2.matrix[0]) {
		return errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] + m2.matrix[r][c]
		}
	}

	return nil
}

// Subtract will subtract the values of a matrix from this matrix.
func (m *MutableMatrix) Subtract(m2 *MutableMatrix) error {
	if len(m.matrix) != len(m2.matrix) {
		return errors.New("width of both matrices are not the same")
	}

	if len(m.matrix[0]) != len(m2.matrix[0]) {
		return errors.New("height of both matrices are not the same")
	}

	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] - m2.matrix[r][c]
		}
	}

	return nil
}

// ScalarMultiply will multiply this matrix by a given scalar value.
func (m *MutableMatrix) ScalarMultiply(s int) {
	for r := 0; r < len(m.matrix); r++ {
		for c := 0; c < len(m.matrix[r]); c++ {
			m.matrix[r][c] = m.matrix[r][c] * s
		}
	}
}

// Transpose will transpose this matrix.
func (m *MutableMatrix) Transpose() {
	t := make([][]int, len(m.matrix[0]))

	for rt := 0; rt < len(t); rt++ {
		for ct := 0; ct < len(m.matrix); ct++ {
			if t[rt] == nil {
				t[rt] = make([]int, len(m.matrix))
			}
			t[rt][ct] = m.matrix[ct][rt]
		}
	}

	m.matrix = t
}

// MatrixMultiply will multiple the given matrix against this matrix.
func (m *MutableMatrix) MatrixMultiply(m2 *MutableMatrix) error {
	if len(m.matrix[0]) != len(m2.matrix) {
		return errors.New("the dimensions of the matrices are incompatible, try transposing one first")
	}

	height := len(m.matrix)
	width := len(m2.matrix[0])
	n := make([][]int, height)

	for i := 0; i < height; i++ {
		n[i] = make([]int, width)
	}

	for rm := 0; rm < len(m.matrix); rm++ {
		for cm2 := 0; cm2 < len(m2.matrix[rm]); cm2++ {
			product := 0
			for cm := 0; cm < len(m.matrix[rm]); cm++ {
				product = product + m.matrix[rm][cm]*m2.matrix[cm][cm2]
			}
			n[rm][cm2] = product
		}
	}

	m.matrix = n

	return nil
}
