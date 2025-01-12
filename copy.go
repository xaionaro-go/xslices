package xslices

func Clone[T any](src []T) []T {
	result := make([]T, 0, len(src))
	result = append(result, src...)
	return result
}
