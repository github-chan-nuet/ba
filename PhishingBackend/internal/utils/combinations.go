package utils

func Combinations[T comparable](elements []T, maxSize int) [][]T {
	var result [][]T
	var comb []T

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(comb) > 0 && len(comb) <= maxSize {
			// make a copy of comb and append to result
			temp := make([]T, len(comb))
			copy(temp, comb)
			result = append(result, temp)
		}
		if len(comb) == maxSize {
			return
		}

		for i := start; i < len(elements); i++ {
			comb = append(comb, elements[i])
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0)
	return result
}
