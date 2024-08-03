package main

import (
	"fmt"
	"log"

	"github.com/o5h/engine/assets"
	"github.com/o5h/engine/pkg/app"
	"github.com/o5h/engine/pkg/app/input/keyboard"
	"github.com/o5h/engine/pkg/app/input/mouse"
	"github.com/o5h/engine/pkg/core/signal"
	"github.com/o5h/glm"
	"github.com/o5h/glx"
	"github.com/o5h/glx/color"
	"github.com/o5h/glx/light/pointlight"
	"github.com/o5h/glx/material"
	"github.com/o5h/glx/mesh/stl"
)

type Example struct {
	rnd         *glx.Renderer
	cam         *glx.Camera
	mesh        glx.IMesh
	light       *pointlight.Light
	keyboardCon *signal.Connection[keyboard.Event]
	mouseCon    *signal.Connection[mouse.Event]
}

func (example *Example) OnCreate(ctx app.Context) {
	log.Println("Created")
	example.rnd = &glx.Renderer{}
	example.cam = glx.NewCamera()
	example.cam.Viewport = glm.RectI{0, 0, 640, 480}

	example.mesh = assets.MustDecode(stl.Load, "bottle.stl")
	material.Init()
	example.light = &pointlight.Light{
		Color: color.Yellow, Location: glm.Vec3{X: 10, Y: 10, Z: 10}}

	example.keyboardCon = keyboard.Events.Connect(func(e keyboard.Event) {
		log.Println(e)
		if e.Code == keyboard.CodeW {
			example.cam.Pos.Y += 0.1
		}
		if e.Code == keyboard.CodeS {
			example.cam.Pos.Y -= 0.1
		}
		if e.Code == keyboard.CodeD {
			example.cam.Pos.X += 0.1
		}
		if e.Code == keyboard.CodeA {
			example.cam.Pos.X -= 0.1
		}
		fmt.Println(example.cam.Pos)
	})
	// example.mouseCon = ctx.MouseEvents().Connect(func(e mouse.Event) {
	// 	log.Println(e)
	// })

}

func (example *Example) OnUpdate(deltaTime float32) {
	material.SimpleBronze.Apply(example.rnd)
	example.light.Apply(example.rnd)
	example.rnd.StartFrame(example.cam, color.Blue)
	example.rnd.SetModel(glm.Mat4x4_Ident)
	example.mesh.Draw(example.rnd)
	example.rnd.EndFrame()
}

func (example *Example) OnDestroy() {
	log.Println("Application Destroyed")
}

func main() {

	app.Start(&Example{}, app.WithTitle("Hello"))
}
