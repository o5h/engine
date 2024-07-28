package rx

import "slices"

type subject[T any] struct {
	closed    bool
	observers []Observer[T]
}

func NewSubject[T any]() Subject[T] {
	return &subject[T]{}
}

func (s *subject[T]) IsClosed() bool {
	return s.closed
}

func (s *subject[T]) Subscribe(dest Observer[T]) Subscription {
	s.append(dest)
	return NewSubscription(func() {
		s.remove(dest)
	})
}

func (s *subject[T]) append(dest Observer[T]) {
	s.observers = append(s.observers, dest)
}

func (s *subject[T]) remove(dest Observer[T]) {
	s.observers = slices.DeleteFunc(s.observers, func(v Observer[T]) bool {
		return v == dest
	})
}

func (s *subject[T]) Unsubscribe() {

}

func (s *subject[T]) Next(value T) {
	for _, observer := range s.observers {
		observer.Next(value)
	}
}

func (s *subject[T]) Error(err error) {
	for _, observer := range s.observers {
		observer.Error(err)
	}
}

func (s *subject[T]) Complete() {
	for _, observer := range s.observers {
		observer.Complete()
	}
}
