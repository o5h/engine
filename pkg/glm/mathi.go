package glm

func MinMaxI(a, b int32) (int32, int32) {
	if a <= b {
		return a, b
	}
	return b, a
}

func MinI(a, b int32) int32 {
	if a <= b {
		return a
	}
	return b
}

func MaxI(a, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}

func AbsI(i int32) int32 {
	if i <= 0 {
		return -i
	}
	return i
}
