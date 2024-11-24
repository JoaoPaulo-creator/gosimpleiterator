// package gosimpleiterator
package main

func Map[T Object](callbackFn func(T) T, s []T) []T {
	newSlice := make([]T, 0)
	for i := range s {
		currentValue := callbackFn(s[i])
		newSlice = append(newSlice, currentValue)
	}

	return newSlice
}

type Object interface {
	interface{} | string | int | uint8 | float32 | float64
}
