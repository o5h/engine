package rx

type Subscribable[T any] interface {
	Subscribe(Observer[T]) Subscription
}

type Observer[T any] interface {
	Next(T)
	Error(error)
	Complete()
}

type Observable[T any] interface {
	Subscribable[T]
}

type Subscriber[T any] interface {
	Observer[T]
	Subscription
}

type Subscription interface {
	IsClosed() bool
	Unsubscribe()
}

type BehaviorSubject[T any] interface {
	Subject[T]
	Value() T
}

type Subject[T any] interface {
	Observable[T]
	Observer[T]
	Subscription
}
