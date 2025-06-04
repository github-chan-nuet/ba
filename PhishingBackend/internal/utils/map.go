package utils

func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}
