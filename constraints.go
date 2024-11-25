package gosimpleiterator

type Constraint interface {
	interface{} | string | int | uint8 | float32 | float64
}

type ReturnType interface {
	Constraint
}
