package glm

type Frustrum [6]Plane

func (f *Frustrum) SetMatrix(m Mat4x4) {
	f[0].Set(m[3]-m[0], m[7]-m[4], m[11]-m[8], m[15]-m[12])
	f[1].Set(m[3]+m[0], m[7]+m[4], m[11]+m[8], m[15]+m[12])
	f[2].Set(m[3]+m[1], m[7]+m[5], m[11]+m[9], m[15]+m[13])
	f[3].Set(m[3]-m[1], m[7]-m[5], m[11]-m[9], m[15]-m[13])
	f[4].Set(m[3]-m[2], m[7]-m[6], m[11]-m[10], m[15]-m[14])
	f[5].Set(m[3]+m[2], m[7]+m[6], m[11]+m[10], m[15]+m[14])
	f.Normalize()
}

func (f *Frustrum) Normalize() {
	f[0].Normalize()
	f[1].Normalize()
	f[2].Normalize()
	f[3].Normalize()
	f[4].Normalize()
	f[5].Normalize()
}

func (f *Frustrum) IntersectSphere(s *Sphere) bool {
	for i := 0; i < 6; i++ {
		if f[i].DistanceToVec3(&s.Center) > s.Radius {
			return false
		}
	}
	return true
}

func (f *Frustrum) ContainsVec3(k *Vec3) bool {
	for i := 0; i < 6; i++ {
		if f[i].DistanceToVec3(k) < 0 {
			return false
		}
	}
	return true
}
