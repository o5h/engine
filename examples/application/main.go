package main

import (
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/signal"
)

type Example struct {
	keyboardCon *signal.Connection[keyboard.Event]
	mouseCon    *signal.Connection[mouse.Event]
}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")

	example.keyboardCon = ctx.KeyboardEvents().Connect(func(e keyboard.Event) {
		log.Println(e)
		if e.Code == keyboard.Code1 {
			example.mouseCon.Disconnect()
		}
	})
	example.mouseCon = ctx.MouseEvents().Connect(func(e mouse.Event) {
		log.Println(e)
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
