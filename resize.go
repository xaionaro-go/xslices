package xslices

import "fmt"

func Resize[T any](in []T, newSize, newMinCap int) []T {
	if newSize > newMinCap {
		panic(fmt.Errorf("the selected size is greater than the selected capacity: %d > %d", newSize, newMinCap))
	}
	if cap(in) >= newMinCap {
		return in[:newSize]
	}
	out := make([]T, newSize, newMinCap)
	copy(out, in)
	return out
}
