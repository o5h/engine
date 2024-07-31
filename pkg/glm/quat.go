package glm

type Quat struct {
	X, Y, Z, W float32
}

func ZeroQuat() *Quat { return &Quat{0, 0, 0, 1} }

func (q *Quat) ToEuler() (yaw, pitch, roll float32) {
	sqw := q.W * q.W
	sqx := q.X * q.X
	sqy := q.Y * q.Y
	sqz := q.Z * q.Z
	unit := sqx + sqy + sqz + sqw // if normalised is one, otherwise is correction factor
	test := q.X*q.Y + q.Z*q.W
	if test > 0.499*unit { // singularity at north pole
		yaw = 2 * Atan2(q.X, q.W)
		pitch = Pi * 0.5
		roll = 0
		return
	}
	if test < -0.499*unit { // singularity at south pole
		yaw = -2 * Atan2(q.X, q.W)
		pitch = -Pi * 0.5
		roll = 0
		return
	}
	yaw = Atan2(2*q.Y*q.W-2*q.X*q.Z, sqx-sqy-sqz+sqw)
	pitch = Asin(2 * test / unit)
	roll = Atan2(2*q.X*q.W-2*q.Y*q.Z, -sqx+sqy-sqz+sqw)
	return
}

func (q *Quat) FromEulerXYZ(x, y, z float32) {
	sinYaw, cosYaw := SinCos(x * 0.5)
	s2, c2 := SinCos(y * 0.5)
	s3, c3 := SinCos(z * 0.5)

	c1c2 := cosYaw * c2
	s1s2 := sinYaw * s2

	q.W = c1c2*c3 - s1s2*s3
	q.X = c1c2*s3 + s1s2*c3
	q.Y = sinYaw*c2*c3 + cosYaw*s2*s3
	q.Z = cosYaw*s2*c3 - sinYaw*c2*s3
}

func (q *Quat) SetEulerXYZ(x, y, z float32) {
	sZ, cZ := SinCos(z * 0.5)
	sY, cY := SinCos(y * 0.5)
	sX, cX := SinCos(x * 0.5)
	cXsY := cX * sY
	sXcY := sX * cY
	cXcY := cX * cY
	sXSy := sX * sY

	q.X = (cXsY * cZ) + (sXcY * sZ)
	q.Y = (sXcY * cZ) - (cXsY * sZ)
	q.Z = (cXcY * sZ) - (sXSy * cZ)
	q.W = (cXcY * cZ) + (sXSy * sZ)
}

func (q *Quat) FromEuler(v *Vec3) {
	q.FromEulerXYZ(v.X, v.Y, v.Z)
}

func (q *Quat) SetZeroRotation() {
	q.X = 0
	q.Y = 0
	q.Z = 0
	q.W = 1
}

func (q *Quat) Set(k *Quat) {
	q.X = k.X
	q.Y = k.Y
	q.Z = k.Z
	q.W = k.W
}

func (q *Quat) SetXYZW(x, y, z, w float32) {
	q.X = x
	q.Y = y
	q.Z = z
	q.W = w
}

func (q *Quat) SetQuat(a *Quat) {
	q.X = a.X
	q.Y = a.Y
	q.Z = a.Z
	q.W = a.W
}

func (q *Quat) SetSum(a, b *Quat) {
	q.X = a.X + b.X
	q.Y = a.Y + b.Y
	q.Z = a.Z + b.Z
	q.W = a.W + b.W
}

func (q *Quat) Normalize() {
	f := 1.0 / Sqrt(q.X*q.X+q.Y*q.Y+q.Z*q.Z+q.W*q.W)
	q.X *= f
	q.Y *= f
	q.Z *= f
	q.W *= f
}

func (q *Quat) Mul(k *Quat) {
	x := q.W*k.X + q.X*k.W + q.Y*k.Z - q.Z*k.Y
	y := q.W*k.Y - q.X*k.Z + q.Y*k.W + q.Z*k.X
	z := q.W*k.Z + q.X*k.Y - q.Y*k.X + q.Z*k.W
	w := q.W*k.W - q.X*k.X - q.Y*k.Y - q.Z*k.Z
	q.X = x
	q.Y = y
	q.Z = z
	q.W = w
}

func (q *Quat) SetMul(v, k *Quat) {
	x := v.W*k.X + v.X*k.W + v.Y*k.Z - v.Z*k.Y
	y := v.W*k.Y - v.X*k.Z + v.Y*k.W + v.Z*k.X
	z := v.W*k.Z + v.X*k.Y - v.Y*k.X + v.Z*k.W
	w := v.W*k.W - v.X*k.X - v.Y*k.Y - v.Z*k.Z
	q.X = x
	q.Y = y
	q.Z = z
	q.W = w
}
