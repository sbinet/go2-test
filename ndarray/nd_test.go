package ndarray_test

import (
	"ndarray"
	"testing"
)

func TestArray1D(t *testing.T) {
	shape := []int{10}
	arr := ndarray.New(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestArray2D(t *testing.T) {
	shape := []int{2, 5}
	arr := ndarray.New(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
