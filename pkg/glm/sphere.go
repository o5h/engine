package glm

import "fmt"

type Sphere struct {
	Center Vec3
	Radius float32
}

func NewSphere(v *Vec3, r float32) *Sphere { return &Sphere{*v, r} }

func (s *Sphere) CylinderCoord(n *Vec3) *Vec2 {
	a := Atan2(n.X, n.Y)
	return &Vec2{a / Pi2, (1 - n.X) / 2}
}

func (s *Sphere) Midpoint(a, b *Vec3) (mid *Vec3) {
	mid = Vec3Add(a, b)
	mid.Normalize()
	mid.Scale(s.Radius)
	return
}

func (s *Sphere) DistanceToVec3(k *Vec3) float32 {
	return s.Center.DistanceTo(k) - s.Radius
}

func (s *Sphere) ContainsVec3(k *Vec3) bool {
	return s.Center.DistanceToSquared(k) <= s.Radius*s.Radius
}

func (s *Sphere) IntersectSphere(o *Sphere) bool {
	ls := LengthSquared(s.Center.X-o.Center.X, s.Center.Y-o.Center.Y, s.Center.Z-o.Center.Z)
	rs := s.Radius*s.Radius + o.Radius*o.Radius
	return ls < rs
}

func (s *Sphere) String() string {
	return fmt.Sprintf("%v %v", s.Center, s.Radius)
}
