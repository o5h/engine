package rx

import (
	"log"
	"testing"
)

func TestExample(t *testing.T) {
	observable := NewObservable(func(s Observer[string]) {
		s.Next("Hello")
	})
	sub := observable.Subscribe(NewObserver(func(s string) {
		log.Println(s)
	}, nil, nil))
	sub.Unsubscribe()

}
