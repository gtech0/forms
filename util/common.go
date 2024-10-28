package util

func Zero[T any]() T {
	return *new(T)
}

func FindKeyByValue(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
