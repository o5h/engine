package rx

func NewSubscription(complete func()) Subscription {
	return &subscription{
		closed:   false,
		complete: complete}
}

type subscription struct {
	closed   bool
	complete func()
}

func (sub *subscription) IsClosed() bool {
	return sub.closed
}

func (sub *subscription) Unsubscribe() {
	if sub.closed {
		return
	}
	if sub.complete == nil {
		return
	}
	sub.complete()
	sub.closed = true
}
