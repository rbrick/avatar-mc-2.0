package main

type MatrixOp func(matrixValue float64) float64

func Add(additiveNumber float64) MatrixOp {
	return func(matrixValue float64) float64 {
		return matrixValue + additiveNumber
	}
}

func Sub(subtractiveNumber float64) MatrixOp {
	return func(matrixValue float64) float64 {
		return matrixValue - subtractiveNumber
	}
}

func Scalar(multiplicativeNumber float64) MatrixOp {
	return func(matrixValue float64) float64 {
		return matrixValue * multiplicativeNumber
	}
}

type Matrix struct {
	size         int
	backingArray []float64
}

func (m *Matrix) Set(y, x int, v float64) {
	// the matrix is stored linearly
	// index = y * size + x
	m.backingArray[y*m.size+x] = v
}

// Add a matrix to this one
func (m *Matrix) Add(other *Matrix) *Matrix {
	dest := Create(m.size)

	for idx, f := range other.backingArray {
		m.executeMatrixOp(dest, idx, Add(f))
	}
	return dest
}

func (m *Matrix) Subtract(other *Matrix) *Matrix {
	dest := Create(m.size)

	for idx, f := range other.backingArray {
		m.executeMatrixOp(dest, idx, Sub(f))
	}
	return dest
}

// Scale executes a scalar function on this matrix
func (m *Matrix) Scale(by float64) {
	for idx := range m.backingArray {
		m.executeMatrixOp(m, idx, Scalar(by))
	}
}

func (m *Matrix) executeMatrixOp(dest *Matrix, idx int, op MatrixOp) *Matrix {
	result := op(m.backingArray[idx])
	dest.backingArray[idx] = result
	return dest
}

// Creates a square matrix (size x size)
func Create(size int) *Matrix {
	return &Matrix{
		size:         size, // internal cache of the size
		backingArray: make([]float64, size*size),
	}
}
