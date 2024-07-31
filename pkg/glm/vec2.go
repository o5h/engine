package glm

type Vec2 struct{ X, Y float32 }

func (v *Vec2) SetSub(a, b *Vec2) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
}

func (v *Vec2) SetXY(x, y float32) {
	v.X = x
	v.Y = y
}

func (v *Vec2) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
}

func (v *Vec2) Length() float32 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) DistanceTo(k *Vec2) float32 {
	return LengthXY(v.X-k.X, v.Y-k.Y)
}

func Vec2Equals(v1, v2 *Vec2) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	if v1.DistanceTo(v2) < Epsilon {
		return true
	}
	return false
}

func Vec2Sub(v1, v2 *Vec2) *Vec2 { return &Vec2{v1.X - v2.X, v1.Y - v2.Y} }
