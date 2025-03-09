package utils

type Pair[T, U any] struct {
	First  T
	Second U
}

func NewPair[T, U any](first T, second U) *Pair[T, U] {
	return &Pair[T, U]{First: first, Second: second}
}

type Triple[T, U, V any] struct {
	First  T
	Second U
	Third  V
}

func NewTriple[T, U, V any](first T, second U, third V) *Triple[T, U, V] {
	return &Triple[T, U, V]{First: first, Second: second, Third: third}
}
