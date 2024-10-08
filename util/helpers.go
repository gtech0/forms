package util

func ConcatSlice[T any](target []T, source ...T) {
	for _, element := range source {
		target = append(target, element)
	}
}
