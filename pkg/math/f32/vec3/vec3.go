package vec3

import (
	"github.com/o5h/engine/pkg/math/f32/math"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (v *Vec3) SetSum(a, b *Vec3) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
}

func (v *Vec3) SetSub(a, b *Vec3) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
}

func (v *Vec3) SetMul(a, b *Vec3) {
	v.X = a.X * b.X
	v.Y = a.Y * b.Y
	v.Z = a.Z * b.Z
}

func (v *Vec3) Scale(factor float32) {
	v.X *= factor
	v.Y *= factor
	v.Z *= factor
}

func (v *Vec3) DistanceTo(k *Vec3) float32 {
	return math.LengthXYZ(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) DistanceToSquared(k *Vec3) float32 {
	return math.LengthSquared(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
	v.Z *= f
}
