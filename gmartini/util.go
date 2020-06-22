package gmartini

func maxUint16(vars ...uint16) uint16 {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func minUint16(vars ...uint16) uint16 {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func maxFloat32(vars ...float32) float32 {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func minFloat32(vars ...float32) float32 {
	min := vars[0]
	for _, v := range vars {
		if v < min {
			min = v
		}
	}
	return min
}

func absInt32(n int32) int32 {
	y := n >> 31
	return (n ^ y) - y
}

func absFloat32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func maxminFloat32(arr []float32) (float32, float32) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func equalFloat32(a, b []float32, tol float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if absFloat32((v - b[i])) > tol {
			return false
		}
	}
	return true
}

func equalInt32(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
