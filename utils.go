package sumcomp

import "math/rand"

func Pick[T any](n int, xs []T) []T {
	ys := make([]T, 0)
	for i := 0; i < n; i++ {
		ys = append(ys, xs[rand.Intn(len(xs))])
	}
	return ys
}
