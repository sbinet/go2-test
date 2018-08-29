package ndarray

type T float64

type Array struct {
	data  []T
	shape []int
}

func New(shape []int, data T) Array {
	return Array{
		data:  data,
		shape: shape,
	}
}

func Sum(arr Array, axis []int) T {
	var t T
	return t
}
