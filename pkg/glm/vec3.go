package glm

import (
	"encoding/json"
	"fmt"
)

type Vec3 struct {
	X, Y, Z float32
}

func Vec3Add(v1, v2 *Vec3) *Vec3 {
	return &Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Vec3Sub(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v *Vec3) Set(a *Vec3) {
	v.X = a.X
	v.Y = a.Y
	v.Z = a.Z
}

func (v *Vec3) SetXYZ(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v *Vec3) SetMin(x, y, z float32) {
	v.X = Min(v.X, x)
	v.Y = Min(v.Y, x)
	v.Z = Min(v.Z, x)
}

func (v *Vec3) SetMax(x, y, z float32) {
	v.X = Max(v.X, x)
	v.Y = Max(v.Y, x)
	v.Z = Max(v.Z, x)
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

func (v *Vec3) AddXYZ(x, y, z float32) {
	v.X += x
	v.Y += y
	v.Z += z
}

func (v *Vec3) Sub(k Vec3) {
	v.X -= k.X
	v.Y -= k.Y
	v.Z -= k.Z
}

func (v *Vec3) Add(k Vec3) {
	v.X += k.X
	v.Y += k.Y
	v.Z += k.Z
}

func (v *Vec3) AddVec3(k *Vec3) {
	v.X += k.X
	v.Y += k.Y
	v.Z += k.Z
}

func (v *Vec3) SetSubVec3(k *Vec3) {
	v.X -= k.X
	v.Y -= k.Y
	v.Z -= k.Z
}

func (v *Vec3) DistanceTo(k *Vec3) float32 {
	return LengthXYZ(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) DistanceToSquared(k *Vec3) float32 {
	return LengthSquared(v.X-k.X, v.Y-k.Y, v.Z-k.Z)
}

func (v *Vec3) Length() float32 {
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Normalize() {
	f := 1.0 / v.Length()
	v.X *= f
	v.Y *= f
	v.Z *= f
}

func (v *Vec3) ApplyMat4x4(m *Mat4x4) {
	x := v.X
	y := v.Y
	z := v.Z

	v.X = x*m[0] + y*m[4] + z*m[8]
	v.Y = x*m[1] + y*m[5] + z*m[9]
	v.Z = x*m[2] + y*m[6] + z*m[10]
}

func (v *Vec3) ApplyQuat(q *Quat) {
	x := v.X
	y := v.Y
	z := v.Z

	qx := q.X
	qy := q.Y
	qz := q.Z
	qw := q.W

	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := -qx*x - qy*y - qz*z

	v.X = ix*qw + iw*-qx + iy*-qz - iz*-qy
	v.Y = iy*qw + iw*-qy + iz*-qx - ix*-qz
	v.Z = iz*qw + iw*-qz + ix*-qy - iy*-qx
}

func (v *Vec3) ApplyMat3x3(m *Mat3x3) {
	x := v.X
	y := v.Y
	z := v.Z

	v.X = m[0]*x + m[3]*y + m[6]*z
	v.Y = m[1]*x + m[4]*y + m[7]*z
	v.Z = m[2]*x + m[5]*y + m[8]*z
}

func (v *Vec3) String() string {
	return fmt.Sprintf("[%f,%f,%f]", v.X, v.Y, v.Z)
}

func (v *Vec3) MarshalJSON() ([]byte, error) {
	return json.Marshal([3]float32{v.X, v.Y, v.Z})
}

func (v *Vec3) UnmarshalJSON(data []byte) error {
	aux := make([]float32, 3)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = aux[0]
	v.Y = aux[1]
	v.Z = aux[2]
	return nil
}

func Vec3Equals(v1, v2 *Vec3) bool {
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

func (v *Vec3) EqEpsilon(k Vec3) bool   { return v.Distance(k) < Epsilon }
func (v *Vec3) Distance(k Vec3) float32 { return LengthXYZ(v.X-k.X, v.Y-k.Y, v.Z-k.Z) }
