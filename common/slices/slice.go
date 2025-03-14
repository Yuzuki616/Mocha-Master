package slices

func TypeTo[T, T2 any](s []T, f func(t T) T2) []T2 {
	temp := make([]T2, len(s))
	for i, v := range s {
		temp[i] = f(v)
	}
	return temp
}
