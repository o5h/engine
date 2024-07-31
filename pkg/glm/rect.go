package glm

import "fmt"

type Rect struct {
	X, Y, W, H float32
}

type RectI struct {
	X, Y, W, H int32
}

func (r *Rect) Set(x, y, w, h float32) {
	r.X = x
	r.Y = y
	r.W = w
	r.H = h
}

func (r *Rect) Aspect() float32 {
	return r.W / r.H
}

func (r *Rect) String() string {
	return fmt.Sprintf("[%f, %f, %f, %f]", r.X, r.Y, r.W, r.H)
}
