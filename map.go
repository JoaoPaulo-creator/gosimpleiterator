package gosimpleiterator

func Map[T Constraint, R ReturnType](callbackFn func(T) R, s []T) []R {
	newSlice := make([]R, 0)
	for i := range s {
		currentValue := callbackFn(s[i])
		newSlice = append(newSlice, currentValue)
	}

	return newSlice
}

type Constraint interface {
	interface{} | string | int | uint8 | float32 | float64
}

type ReturnType interface {
	Constraint
}
