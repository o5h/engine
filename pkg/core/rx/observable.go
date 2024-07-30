package rx

type Observable[T any] struct {
	observer func(Observer[T])
}

type subscriber[T any] struct {
	Observable[T]
	Subscription
}

func NewObservable[T any](fn func(Observer[T])) *Observable[T] {
	return &Observable[T]{observer: fn}
}

func (o *Observable[T]) Subscribe(observer Observer[T]) *Subscription {
	subscription := newSubscription(func() {})
	go o.observer(observer)
	return subscription
}
