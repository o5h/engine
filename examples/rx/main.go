package main

import (
	"log"
	"time"

	"github.com/o5h/engine/pkg/core/rx"
)

func main() {
	observable := rx.NewObservable(func(s rx.Observer[string]) {
		s.Next("Hello")
		time.Sleep(3 * time.Second)
		s.Next("World")
		time.Sleep(1 * time.Second)
		s.Complete()
	})
	sub := observable.Subscribe(rx.NewObserver(func(s string) {
		log.Println(s)
	}, nil, func() {
		log.Println("Complete")
	}))
	time.Sleep(2 * time.Second)
	sub.Unsubscribe()
	time.Sleep(4 * time.Second)
}
