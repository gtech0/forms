package util

func Zero[T any]() T {
	return *new(T)
}
