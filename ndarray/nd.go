package ndarray

contract addable(x T) {
	x += x
}

type Array(type T) struct {
	data  []T
	shape []int
}

func New(type T)(shape []int, data []T) Array(T) {
	return Array(T){
		data:  data,
		shape: shape,
	}
}

func Sum(type T addable)(arr Array(T), axis []int) T {
	var sum T
	for _, v := range arr.data {
		sum += v
	}
	return sum
}
