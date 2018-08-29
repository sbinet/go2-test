package ndarray

contract Elem(x T) {
	x += x
	x = x
	x == x
	x -= x
}

type Array(type T Elem) struct {
	data  []T
	shape []int
}

func New(type T Elem)(shape []int, data []T) Array(T) {
	return Array(T){
		data:  data,
		shape: shape,
	}
}

func Sum(type T Elem)(arr Array(T), axis []int) T {
	var sum T
	for _, v := range arr.data {
		sum += v
	}
	return sum
}
