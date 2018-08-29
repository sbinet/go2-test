package ndarray

type Array(type T) struct {
	data  []T
	shape []int
}

func New(type T)(shape []int, data []T) Array {
	return Array{
		data:  data,
		shape: shape,
	}
}

func Sum(type T)(arr Array, axis []int) T {
	var sum T
	for _, v := range arr.data {
		sum += v
	}
	return sum
}
