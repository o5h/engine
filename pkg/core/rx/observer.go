package rx

type Observer[T any] func(T, bool, error)
type observer[T any] struct{ Observer[T] }
