package rx

import "slices"

type Subject[T any] struct {
	observers []*observer[T]
}

func (s *Subject[T]) OnNext(fn func(T)) *Subscription {
	return s.Subscribe(func(t T, b bool, err error) { fn(t) })
}

func (s *Subject[T]) Subscribe(dest Observer[T]) *Subscription {
	observer := &observer[T]{Observer: dest}
	s.append(observer)
	return newSubscription(func() {
		s.remove(observer)
	})
}

func (s *Subject[T]) Next(v T, complete bool, err error) {
	for _, o := range s.observers {
		o.Observer(v, complete, err)
	}
}

func (s *Subject[T]) append(dest *observer[T]) {
	s.observers = append(s.observers, dest)
}

func (s *Subject[T]) remove(dest *observer[T]) {
	s.observers = slices.DeleteFunc(s.observers, func(v *observer[T]) bool {
		return v == dest
	})
}
