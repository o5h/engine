package rx

func Pipe1[A any, B any](ab Operator[A, B]) {
	//???
}

type Pipe[A any, B any] struct {
	A  Observable[A]
	AB Operator[A, B]
}

func (pipe Pipe[A, B]) Subscribe(observer Observer[B]) Subscription {
	return pipe.AB(pipe.A).Subscribe(observer)
}

type Operator[A any, B any] func(Observable[A]) Observable[B]

func Map[A any, B any](fn func(A) B) Operator[A, B] {
	return func(source Observable[A]) Observable[B] {
		return &observableMap[A, B]{fn: fn, source: source}
	}
}

type observableMap[A any, B any] struct {
	fn     func(A) B
	source Observable[A]
}

func (o *observableMap[A, B]) Subscribe(dest Observer[B]) Subscription {
	subscription := o.source.Subscribe(NewObserver(
		func(value A) { dest.Next(o.fn(value)) },
		func(err error) { dest.Error(err) },
		func() { dest.Complete() }))
	subscriber := NewSubscriber(dest, NewSubscription(func() {
		subscription.Unsubscribe()
	}))
	return subscriber
}
