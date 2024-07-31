package glm

// 0  3  6
// 1  4  7
// 2  5  8

type Mat3x3 [9]float32

var Mat3x3_Ident = Mat3x3{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1}

func (m *Mat3x3) SetIdentity() {
	m[0] = 1
	m[1] = 0
	m[2] = 0

	m[3] = 0
	m[4] = 1
	m[5] = 0

	m[6] = 0
	m[7] = 0
	m[8] = 1
}

func (m *Mat3x3) FromQuaternion(q *Quat) {
	xx := q.X * q.X
	xy := q.X * q.Y
	xz := q.X * q.Z
	xw := q.X * q.W
	yy := q.Y * q.Y
	yz := q.Y * q.Z
	yw := q.Y * q.W
	zz := q.Z * q.Z
	zw := q.Z * q.W

	// Compute rotation tranformation
	m[0] = 1 - 2*(yy+zz)
	m[1] = 2 * (xy - zw)
	m[2] = 2 * (xz + yw)

	m[3] = 2 * (xy + zw)
	m[4] = 1 - 2*(xx+zz)
	m[5] = 2 * (yz - xw)

	m[6] = 2 * (xz - yw)
	m[7] = 2 * (yz + xw)
	m[8] = 1 - 2*(xx+yy)
}

func (m *Mat3x3) SetFormEulerXYZ(angleX, angleY, angleZ float32) {
	sinX, cosX := SinCos(angleX)
	sinY, cosY := SinCos(angleY)
	sinZ, cosZ := SinCos(angleZ)

	mSinX := -sinX
	mSinY := -sinY
	mSinZ := -sinZ

	mSinXTomSinY := mSinX * mSinY
	cosXTomMinY := cosX * mSinY

	m[0] = cosY * cosZ
	m[1] = mSinXTomSinY*cosZ + cosX*sinZ
	m[2] = cosXTomMinY*cosZ + sinX*sinZ
	m[3] = cosY * mSinZ
	m[4] = mSinXTomSinY*mSinZ + cosX*cosZ
	m[5] = cosXTomMinY*mSinZ + sinX*cosZ
	m[6] = sinY
	m[7] = mSinX * cosY
	m[8] = cosX * cosY

}
