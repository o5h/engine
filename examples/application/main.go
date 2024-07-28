package main

import (
	"fmt"
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/rx"
)

type Example struct{}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")
	ctx.MouseEvents().Subscribe(rx.NewObserver[mouse.Event](
		func(e mouse.Event) {
			fmt.Println(e)
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
