package gosimpleiterator

func Filter[T Constraint](preficate func(T) bool, s []T) []T {
	newSlice := make([]T, 0)
	for i := range s {
		if preficate(s[i]) {
			newSlice = append(newSlice, s[i])
		}
	}

	return newSlice
}
