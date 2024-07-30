package rx

type observer[T any] struct {
	next     func(T)
	err      func(error)
	complete func()
}

func NewObserver[T any](next func(T),
	err func(error),
	complete func()) Observer[T] {
	return &observer[T]{
		next:     next,
		err:      err,
		complete: complete}
}

func (o *observer[T]) Next(v T) {
	if o.next != nil {
		o.next(v)
	}
}

func (o *observer[T]) Error(err error) {
	if o.err != nil {
		o.err(err)
	}
}

func (o *observer[T]) Complete() {
	if o.complete != nil {
		o.complete()
	}
}
