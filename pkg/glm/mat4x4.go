package glm

import (
	"errors"
	"fmt"
	"unsafe"
)

//  | 0  4  8  12 |
//  | 1  5  9  13 |
//  | 2  6  10 14 |
//  | 3  7  11 15 |

type Mat4x4 [16]float32

var (
	Ident = Mat4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
)

func (m *Mat4x4) Ptr() uintptr {
	return uintptr(unsafe.Pointer(&m[0]))
}
func (m Mat4x4) SetIdentity() {
	m[0] = 1
	m[1] = 0
	m[2] = 0
	m[3] = 0

	m[4] = 0
	m[5] = 1
	m[6] = 0
	m[7] = 0

	m[8] = 0
	m[9] = 0
	m[10] = 1
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

func (m *Mat4x4) LookAt(eye, target, up Vec3) {

	fx := target.X - eye.X
	fy := target.Y - eye.Y
	fz := target.Z - eye.Z

	// Normalize f
	rlf := 1.0 / LengthXYZ(fx, fy, fz)
	fx *= rlf
	fy *= rlf
	fz *= rlf

	// compute s = f x up (x means "cross product")
	sx := fy*up.Z - fz*up.Y
	sy := fz*up.X - fx*up.Z
	sz := fx*up.Y - fy*up.X

	// and normalize s
	rls := 1.0 / LengthXYZ(sx, sy, sz)
	sx *= rls
	sy *= rls
	sz *= rls

	// compute u = s x f
	ux := sy*fz - sz*fy
	uy := sz*fx - sx*fz
	uz := sx*fy - sy*fx

	m[0] = sx
	m[1] = ux
	m[2] = -fx
	m[3] = 0.0

	m[4] = sy
	m[5] = uy
	m[6] = -fy
	m[7] = 0.0

	m[8] = sz
	m[9] = uz
	m[10] = -fz
	m[11] = 0.0

	// m[12] = 0.0
	// m[13] = 0.0
	// m[14] = 0.0
	// m[15] = 1.0

	//m.Translate(-eyeX, -eyeY, -eyeZ)
	m[12] = m[0]*-eye.X + m[4]*-eye.Y + m[8]*-eye.Z
	m[13] = m[1]*-eye.X + m[5]*-eye.Y + m[9]*-eye.Z
	m[14] = m[2]*-eye.X + m[6]*-eye.Y + m[10]*-eye.Z
	m[15] = 1.0
}

func (m *Mat4x4) TranslateXYZ(x, y, z float32) {
	m[12] += m[0]*x + m[4]*y + m[8]*z
	m[13] += m[1]*x + m[5]*y + m[9]*z
	m[14] += m[2]*x + m[6]*y + m[10]*z
	m[15] += m[3]*x + m[7]*y + m[11]*z
}

func (m *Mat4x4) Perspective(zNear, zFar, fovy, aspect float32) {
	f := 1.0 / Tan(fovy*(Pi/360.0))
	rangeReciprocal := 1.0 / (zNear - zFar)

	m[0] = f / aspect
	m[1] = 0.0
	m[2] = 0.0
	m[3] = 0.0

	m[4] = 0.0
	m[5] = f
	m[6] = 0.0
	m[7] = 0.0

	m[8] = 0.0
	m[9] = 0.0
	m[10] = (zFar + zNear) * rangeReciprocal
	m[11] = -1.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 2.0 * zFar * zNear * rangeReciprocal
	m[15] = 0.0
}

func (m *Mat4x4) Frustum(left, right, bottom, top, near, far float32) {
	width := 1.0 / (right - left)
	height := 1.0 / (top - bottom)
	depth := 1.0 / (near - far)

	m[0] = 2.0 * (near * width)
	m[1] = 0.0
	m[2] = 0.0
	m[3] = 0.0

	m[4] = 0.0
	m[5] = 2.0 * (near * height)
	m[6] = 0.0
	m[7] = 0.0

	m[8] = (right + left) * width
	m[9] = (top + bottom) * height
	m[10] = (far + near) * depth
	m[11] = -1.0

	m[12] = 0.0
	m[13] = 0.0
	m[14] = 2.0 * (far * near * depth)
	m[15] = 0.0
}

func (m *Mat4x4) Ortho(left, right, bottom, top, near, far float32) {

	m[0] = 2 / (right - left)
	m[1] = 0
	m[2] = 0
	m[3] = 0

	m[4] = 0
	m[5] = 2 / (top - bottom)
	m[6] = 0
	m[7] = 0

	m[8] = 0
	m[9] = 0
	m[10] = -2 / (far - near)
	m[11] = 0

	m[12] = -(right + left) / (right - left)
	m[13] = -(top + bottom) / (top - bottom)
	m[14] = -(far + near) / (far - near)
	m[15] = 1
}

func (m *Mat4x4) SetRotateX(radians float32) {
	s, c := SinCos(radians)

	m[0] = 1
	m[1] = 0
	m[2] = 0
	m[3] = 0

	m[4] = 0
	m[5] = c
	m[6] = -s
	m[7] = 0

	m[8] = 0
	m[9] = s
	m[10] = c
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

func (m *Mat4x4) SetRotateY(radians float32) {
	s, c := SinCos(radians)

	m[0] = c
	m[1] = 0
	m[2] = s
	m[3] = 0

	m[4] = 0
	m[5] = 1
	m[6] = 0
	m[7] = 0

	m[8] = -s
	m[9] = 0
	m[10] = c
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

func (m *Mat4x4) SetRotateZ(radians float32) {
	s, c := SinCos(radians)

	m[0] = c
	m[1] = -s
	m[2] = 0
	m[3] = 0

	m[4] = s
	m[5] = c
	m[6] = 0
	m[7] = 0

	m[8] = 0
	m[9] = 0
	m[10] = 1
	m[11] = 0

	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

// Transform Order is : Scale, Rotate, Translate
func (m *Mat4x4) SetTransform(t *Transform) {
	rot := Mat3x3{}
	rot.SetFormEulerXYZ(t.Rotation.X, t.Rotation.Y, t.Rotation.Z)

	m[0] = t.Scale.X * rot[0]
	m[1] = t.Scale.X * rot[1]
	m[2] = t.Scale.X * rot[2]
	m[3] = 0

	m[4] = t.Scale.Y * rot[3]
	m[5] = t.Scale.Y * rot[4]
	m[6] = t.Scale.Y * rot[5]
	m[7] = 0

	m[8] = t.Scale.Z * rot[6]
	m[9] = t.Scale.Z * rot[7]
	m[10] = t.Scale.Z * rot[8]
	m[11] = 0

	m[12] = t.Position.X
	m[13] = t.Position.Y
	m[14] = t.Position.Z
	m[15] = 1

}

func (m *Mat4x4) Transform2(location *Vec3, scale *Vec3, orient *Quat) {
	rot := Mat3x3{}
	rot.FromQuaternion(orient)

	// Ordering:
	//    1. Scale
	//    2. Rotate
	//    3. Translate

	// Set up final matrix with scale, rotation and translation
	m[0] = scale.X * rot[0]
	m[1] = scale.Y * rot[1]
	m[2] = scale.Z * rot[2]
	m[3] = location.X

	m[4] = scale.X * rot[3]
	m[5] = scale.Y * rot[4]
	m[6] = scale.Z * rot[5]
	m[7] = location.Y

	m[8] = scale.X * rot[6]
	m[9] = scale.Y * rot[7]
	m[10] = scale.Z * rot[8]
	m[11] = location.Z

	// No projection term
	m[12] = 0
	m[13] = 0
	m[14] = 0
	m[15] = 1
}

func (m *Mat4x4) SetMul(m1, m2 Mat4x4) {
	m[0] = m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3]
	m[1] = m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3]
	m[2] = m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3]
	m[3] = m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]

	m[4] = m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7]
	m[5] = m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7]
	m[6] = m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7]
	m[7] = m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]

	m[8] = m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11]
	m[9] = m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11]
	m[10] = m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11]
	m[11] = m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]

	m[12] = m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15]
	m[13] = m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15]
	m[14] = m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15]
	m[15] = m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]
}

func (m *Mat4x4) Transpose() {
	m[1], m[4] = m[4], m[1]
	m[2], m[8] = m[8], m[2]
	m[6], m[9] = m[9], m[6]
	m[3], m[12] = m[12], m[3]
	m[7], m[13] = m[13], m[7]
	m[11], m[14] = m[14], m[11]
}

func (m *Mat4x4) String() string {
	return fmt.Sprintf(
		"{\n[%f, %f, %f, %f]\n[%f, %f, %f, %f]\n[%f, %f, %f, %f]\n[%f, %f, %f, %f]\n}",
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15])
}

func (m *Mat4x4) CopyInverseFrom(src *Mat4x4) error {
	a00 := src[0]
	a01 := src[1]
	a02 := src[2]
	a03 := src[3]
	a10 := src[4]
	a11 := src[5]
	a12 := src[6]
	a13 := src[7]
	a20 := src[8]
	a21 := src[9]
	a22 := src[10]
	a23 := src[11]
	a30 := src[12]
	a31 := src[13]
	a32 := src[14]
	a33 := src[15]

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	det := b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06
	if det == 0.0 {
		*m = *src
		return errors.New("copyInverseFrom: null determinant")
	}
	invDet := 1.0 / det

	m[0] = (a11*b11 - a12*b10 + a13*b09) * invDet
	m[1] = (-a01*b11 + a02*b10 - a03*b09) * invDet
	m[2] = (a31*b05 - a32*b04 + a33*b03) * invDet
	m[3] = (-a21*b05 + a22*b04 - a23*b03) * invDet
	m[4] = (-a10*b11 + a12*b08 - a13*b07) * invDet
	m[5] = (a00*b11 - a02*b08 + a03*b07) * invDet
	m[6] = (-a30*b05 + a32*b02 - a33*b01) * invDet
	m[7] = (a20*b05 - a22*b02 + a23*b01) * invDet
	m[8] = (a10*b10 - a11*b08 + a13*b06) * invDet
	m[9] = (-a00*b10 + a01*b08 - a03*b06) * invDet
	m[10] = (a30*b04 - a31*b02 + a33*b00) * invDet
	m[11] = (-a20*b04 + a21*b02 - a23*b00) * invDet
	m[12] = (-a10*b09 + a11*b07 - a12*b06) * invDet
	m[13] = (a00*b09 - a01*b07 + a02*b06) * invDet
	m[14] = (-a30*b03 + a31*b01 - a32*b00) * invDet
	m[15] = (a20*b03 - a21*b01 + a22*b00) * invDet

	return nil
}

func (m *Mat4x4) MulVec4(v Vec4) Vec4 {
	return Vec4{
		X: m[0]*v.X + m[4]*v.Y + m[8]*v.Z + m[12]*v.W,
		Y: m[1]*v.X + m[5]*v.Y + m[9]*v.Z + m[13]*v.W,
		Z: m[2]*v.X + m[6]*v.Y + m[10]*v.Z + m[14]*v.W,
		W: m[3]*v.X + m[7]*v.Y + m[11]*v.Z + m[15]*v.W,
	}
}
