package gosimpleiterator

func Map[T Constraint, R ReturnType](callbackFn func(T) R, s []T) []R {
	newSlice := make([]R, 0)
	for i := range s {
		currentValue := callbackFn(s[i])
		newSlice = append(newSlice, currentValue)
	}

	return newSlice
}
