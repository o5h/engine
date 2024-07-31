package glm

type AABB struct {
	Min, Max Vec3
}

func (aabb *AABB) ContainsVec3(k Vec3) bool {
	return aabb.Min.X <= k.X &&
		aabb.Min.Y <= k.Y &&
		aabb.Min.Z <= k.Z &&
		aabb.Max.X >= k.X &&
		aabb.Max.Y >= k.Y &&
		aabb.Max.Z >= k.Z
}

func (aabb *AABB) IntersectAABB(o *AABB) bool {
	return (aabb.Min.X <= o.Max.X && aabb.Max.X >= o.Min.X) &&
		(aabb.Min.Y <= o.Max.Y && aabb.Max.Y >= o.Min.Y) &&
		(aabb.Min.Z <= o.Max.Z && aabb.Max.X >= o.Min.Z)
}

func (aabb *AABB) IntersectSphere(sphere *Sphere) bool {
	x := Clamp(sphere.Center.X, aabb.Min.X, aabb.Max.X)
	y := Clamp(sphere.Center.Y, aabb.Min.Y, aabb.Max.Y)
	z := Clamp(sphere.Center.Z, aabb.Min.Z, aabb.Max.Z)
	ls := LengthSquared(x-sphere.Center.X, y-sphere.Center.Y, z-sphere.Center.Z)
	rs := sphere.Radius * sphere.Radius
	return ls < rs
}
