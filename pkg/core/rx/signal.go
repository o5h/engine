package rx

// import "slices"

// type Subscription[T any] struct {
// 	observable *Observable[T]
// 	callback   func(T)
// }

// func (s *Subscription[T]) Close() {
// 	s.observable.unsubscribe(s)
// 	s.callback = nil
// }

// type Observable[T any] struct {
// 	subscriptions []*Subscription[T]
// }

// func (o *Observable[T]) Next(value T) {
// 	for _, s := range o.subscriptions {
// 		s.callback(value)
// 	}
// }

// func (o *Observable[T]) Subscribe(callback func(T)) *Subscription[T] {
// 	s := &Subscription[T]{observable: o, callback: callback}
// 	o.subscriptions = append(o.subscriptions, s)
// 	return s

// }

// func (o *Observable[T]) unsubscribe(s *Subscription[T]) {
// 	o.subscriptions = slices.DeleteFunc(o.subscriptions, func(ss *Subscription[T]) bool {
// 		return s == ss
// 	})
// }
