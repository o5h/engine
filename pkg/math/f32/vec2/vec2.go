package vec2

var (
	ONE = Vec2{X: 1, Y: 1}
)

type Vec2 struct {
	X float32
	Y float32
}

func (v *Vec2) AddXY(x, y float32) { v.X, v.Y = x, y }

func DotProduct(v1, v2 Vec2) float32 {
	return v1.X*v2.X + v1.Y*v2.Y
}
