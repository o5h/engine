package glm

import (
	"fmt"
)

type Plane struct {
	Normal   Vec3
	Constant float32
}

func (p *Plane) Set(x, y, z, w float32) {
	p.Normal.SetXYZ(x, y, z)
	p.Constant = w
}

func (p *Plane) DistanceToVec3(k *Vec3) float32 {
	return DotProduct(&p.Normal, k) + p.Constant
}

func (p *Plane) Normalize() {
	inverseNormalLength := 1.0 / p.Normal.Length()
	p.Normal.Scale(inverseNormalLength)
	p.Constant *= inverseNormalLength
}

func (p *Plane) String() string { return fmt.Sprintf("[%v, %f]", p.Normal, p.Constant) }

func (p *Plane) IntersectRay(r *Ray) *Vec3 {
	denominator := DotProduct(&p.Normal, &r.Direction)
	if Abs(denominator) > Epsilon {
		difference := p.Normal
		difference.Scale(p.Constant)
		difference.SetSubVec3(&r.Origin)
		t := DotProduct(&difference, &p.Normal) / denominator
		if t > Epsilon {
			intersect := r.Direction
			intersect.Scale(t)
			return Vec3Add(&intersect, &r.Origin)
		}
	}
	return nil
}
