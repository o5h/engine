package math

import "math"

const (
	//Pi float32 Pi
	Pi = float32(math.Pi)
	//Pi2 Pi*2
	Pi2 = float32(math.Pi * 2)
	//Pi180 Pi / 180.0
	Pi180 = float32(math.Pi / 180.0)
	//Epsilon
	Epsilon = 0.000001
)

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

// Deg2Rad converts the number in degrees to the radian equivalent
func Deg2Rad(f float32) float32 {
	return Pi180 * f
}

func SinCos(angle float32) (s, c float32) {
	sin, cos := math.Sincos(float64(angle))
	s = float32(sin)
	c = float32(cos)
	return
}

func Tan(f float32) float32 {
	return float32(math.Tan(float64(f)))
}

func Sqrt(f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

func LengthXY(x, y float32) float32 {
	return Sqrt(x*x + y*y)
}

func LengthXYZ(x, y, z float32) float32 {
	return Sqrt(x*x + y*y + z*z)
}

func Asin(f float32) float32 {
	return float32(math.Asin(float64(f)))
}

func Atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func LengthSquared(x, y, z float32) float32 {
	return x*x + y*y + z*z
}

func Lerp(start, end, t float32) float32 {
	return start*(1-t) + end*t
}

func Clamp(value, min, max float32) float32 {
	return Max(min, Min(value, max))
}

func NextPow2(x uint32) uint32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}
