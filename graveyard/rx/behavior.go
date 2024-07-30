package rx

type behaviorSubject[T any] struct {
	value T
	subject[T]
}

func NewBehaviorSubject[T any]() Subject[T] {
	return &behaviorSubject[T]{}
}

func (s *behaviorSubject[T]) Subscribe(dest Observer[T]) Subscription {
	defer dest.Next(s.value)
	return s.subject.Subscribe(dest)
}
