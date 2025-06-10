package utils

func Find[T any](slice []T, predicate func(T) bool) *T {
	for i := range slice {
		if predicate(slice[i]) {
			return &slice[i]
		}
	}
	return nil
}

func FindAll[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for i := range slice {
		if predicate(slice[i]) {
			result = append(result, slice[i])
		}
	}
	return result
}
