package main

import (
	"log"
	"time"

	"github.com/o5h/engine/pkg/core/rx"
)

func main() {
	observable := rx.NewSubject[string]()
	observable.Next("Hello")
	sub := observable.Subscribe(rx.NewObserver(func(s string) {
		log.Println(s)
	}, nil, func() {
		log.Println("Complete")
	}))
	observable.Next("World")
	time.Sleep(2 * time.Second)
	observable.Complete()
	sub.Unsubscribe()
	observable.Next("!")
	time.Sleep(4 * time.Second)
}
