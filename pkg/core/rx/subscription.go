package rx

type Subscription struct {
	callbacks []func()
}

func (s *Subscription) Unsubscribe() {
	for _, fn := range s.callbacks {
		fn()
	}
	s.callbacks = nil
}

func newSubscription(fn func()) *Subscription {
	return &Subscription{callbacks: []func(){fn}}
}
