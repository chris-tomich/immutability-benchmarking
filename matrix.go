package immutabilitybenchmarking

//const MatrixWidth int = 10
//const MatrixHeight int = 10
//const MatrixWidth int = 30
//const MatrixHeight int = 30
//const MatrixWidth int = 90
//const MatrixHeight int = 90
//const MatrixWidth int = 270
//const MatrixHeight int = 270
const MatrixWidth int = 810
const MatrixHeight int = 810

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
