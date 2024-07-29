package main

import (
	"fmt"
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/rx"
)

type Example struct {
	mouseSubscription rx.Subscription
}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")

	example.mouseSubscription = ctx.MouseEvents().Subscribe(rx.NewObserver(
		func(e mouse.Event) {
			fmt.Println(e)
		}, nil, nil))

	eventToAction := rx.Map(func(e mouse.Event) mouse.Action { return e.Action })
	rx.Pipe[mouse.Event, mouse.Action]{
		A:  ctx.MouseEvents(),
		AB: eventToAction,
	}.Subscribe(rx.NewObserver(func(a mouse.Action) {
		log.Println(a)
	}, nil, nil))

	// rx.Map(func(e mouse.Event) bool { return e.Action == mouse.ActionPress })(ctx.MouseEvents()).
	// Subscribe(func ()  {

	// })

	ctx.KeyboardEvents().Subscribe(rx.NewObserver(
		func(e keyboard.Event) {
			fmt.Println(e)
			if e.Code == keyboard.Escape {
				ctx.Done()
			}
			if e.Code == keyboard.Code1 {
				example.mouseSubscription.Unsubscribe()
			}
		}, nil, nil))
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
