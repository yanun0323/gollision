package gollision

type unsignedInteger interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type signedInteger interface {
	int8 | int16 | int32 | int64 | int
}

type integer interface {
	unsignedInteger | signedInteger
}

func max[T integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func abs[T integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
