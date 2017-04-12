package immutabilitybenchmarking

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
