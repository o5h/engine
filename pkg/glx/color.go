package glx

import (
	"math/rand"

	"github.com/o5h/engine/pkg/glm"
)

type RGB uint32

type RGBA uint32

type Color struct{ R, G, B, A float32 }

var (
	White  = NewColorRGBA(0xFFFFFFFF)
	Black  = NewColorRGBA(0x000000FF)
	Red    = NewColorRGBA(0xFF0000FF)
	Green  = NewColorRGBA(0x00FF00FF)
	Blue   = NewColorRGBA(0x0000FFFF)
	Yellow = NewColorRGBA(0xFFFF00FF)
	Gray   = NewColorRGBA(0x555555FF)
)

func NewColorRGBA(rgba RGBA) Color {
	c := Color{}
	c.ValueOfRGBA(rgba)
	return c
}

func (c *Color) ValueOfRGBA(rgba RGBA) {
	c.R = float32(byte(rgba>>24)) / 255.0
	c.G = float32(byte(rgba>>16)) / 255.0
	c.B = float32(byte(rgba>>8)) / 255.0
	c.A = float32(byte(rgba)) / 255.0
}

func (c *Color) ValueOfRGB(rgb RGB) {
	c.R = float32(byte(rgb>>16)) / 255.0
	c.G = float32(byte(rgb>>8)) / 255.0
	c.B = float32(byte(rgb)) / 255.0
	c.A = 0
}

func RandomRGB() Color {
	return Color{R: rand.Float32(), G: rand.Float32(), B: rand.Float32(), A: 0}
}

func (c *Color) Lerp(o *Color, t float32) {
	c.R = glm.Lerp(c.R, o.R, t)
	c.G = glm.Lerp(c.G, o.G, t)
	c.B = glm.Lerp(c.B, o.B, t)
	c.A = glm.Lerp(c.A, o.A, t)
}
