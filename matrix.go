package immutabilitybenchmarking

const MatrixWidth int = 50
const MatrixHeight int = 50

type Matrix interface {
	Width() int
	Height() int
	Get(int, int) int
	Equals(Matrix) bool
	Add(Matrix) (Matrix, error)
	Subtract(Matrix) (Matrix, error)
	ScalarMultiply(s int) Matrix
	Transpose() Matrix
	MatrixMultiply(Matrix) (Matrix, error)
}
