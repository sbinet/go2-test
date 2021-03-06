package ndarray_test

import (
	"testing"

	"github.com/sbinet/go2-test/ndarray"
)

func TestArray1D(t *testing.T) {
	shape := []int{10}
	arr0 := ndarray.New(float64)(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	arr1 := ndarray.New(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestArray2D(t *testing.T) {
	shape := []int{2, 5}
	arr0 := ndarray.New(float64)(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	arr1 := ndarray.New(shape, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
