package util

func Zero[T any]() T {
	return *new(T)
}

func Pointer[T any](d T) *T {
	return &d
}

func FindKeyByValue(m map[int]int, value int) (key int, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func Difference[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
