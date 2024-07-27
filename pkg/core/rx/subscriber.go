package rx

type subscriber[T any] struct {
	subscription Subscription
	observer     Observer[T]
}

func NewSubscriber[T any](observer Observer[T], subscription Subscription) Subscriber[T] {
	return &subscriber[T]{observer: observer,
		subscription: subscription}

}
func NewSubscriberNextErrorComplete[T any](
	next func(T),
	err func(error),
	complete func()) Subscriber[T] {
	sub := subscriber[T]{}
	sub.observer = NewObserver(next, err, complete)
	sub.subscription = NewSubscription(func() { sub.Complete() })
	return &sub
}

func (sub *subscriber[T]) IsClosed() bool {
	return sub.subscription.IsClosed()
}

func (sub *subscriber[T]) Unsubscribe() {
	sub.subscription.Unsubscribe()
}

func (sub *subscriber[T]) Next(v T) {
	if sub.subscription.IsClosed() {
		return
	}
	sub.observer.Next(v)
}

func (sub *subscriber[T]) Error(err error) {
	if sub.subscription.IsClosed() {
		return
	}
	sub.observer.Error(err)
	sub.subscription.Unsubscribe()
}

func (sub *subscriber[T]) Complete() {
	if sub.subscription.IsClosed() {
		return
	}
	sub.observer.Complete()
	sub.subscription.Unsubscribe()
}
