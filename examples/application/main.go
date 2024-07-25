package main

import (
	"log"

	"github.com/o5h/engine/internal/opengl/gl"
	"github.com/o5h/engine/pkg/app"
)

type Example struct{}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")
}

func (example *Example) OnUpdate(deltaTime float32) {
	gl.ClearColor(1, 0, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (example *Example) OnDestroy() {
	log.Println("Destroyed")
}

func main() {

	app.Start(&Example{})
}
