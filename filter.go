package gosimpleiterator

func Filter[T Constraint](predicate func(T) bool, s []T) []T {
	newSlice := make([]T, 0)
	for i := range s {
		if predicate(s[i]) {
			newSlice = append(newSlice, s[i])
		}
	}

	return newSlice
}
