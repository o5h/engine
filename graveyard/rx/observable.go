package rx

type observable[T any] struct {
	subscribe func(Observer[T])
	ch        chan T
}

func NewObservable[T any](subscribe func(Observer[T])) Observable[T] {
	o := observable[T]{
		subscribe: subscribe,
		ch:        make(chan T, 1)}
	return &o
}

func (o *observable[T]) Subscribe(dest Observer[T]) Subscription {
	subscriber := NewSubscriber(dest, NewSubscription(func() {
		close(o.ch)
	}))
	go o.subscribe(subscriber)
	return subscriber
}
