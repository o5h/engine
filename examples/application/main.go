package main

import (
	"log"
	"time"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/rx"
)

type Example struct {
	keyboardSubscription *rx.Subscription
	mouseSubscription    *rx.Subscription
}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")

	example.keyboardSubscription = ctx.KeyboardEvents().OnNext(func(e keyboard.Event) {
		// log.Println(e)
		if e.Code == keyboard.Code1 {
			example.mouseSubscription.Unsubscribe()
		}
	})
	example.mouseSubscription = ctx.MouseEvents().Subscribe(func(e mouse.Event, b bool, err error) {
		//log.Println(e)
	})

	rx.NewObservable(func(o rx.Observer[string]) {
		o("hello", false, nil)
		time.Sleep(1 * time.Second)
		o("world", false, nil)
		time.Sleep(1 * time.Second)
		o("!!!!!!!!!!!", false, nil)
		o("", true, nil)
	}).Subscribe(func(s string, b bool, err error) {
		log.Println(s)
	})
}

func (example *Example) OnUpdate(deltaTime float32) {
	gl.ClearColor(1, 0, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (example *Example) OnDestroy() {
	log.Println("Destroyed")
}

func main() {

	app.Start(&Example{}, app.WithTitle("Hello"))
}
