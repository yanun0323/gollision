package gollision

type UnsignedInteger interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type SignedInteger interface {
	int8 | int16 | int32 | int64 | int
}

type Integer interface {
	UnsignedInteger | SignedInteger
}

func max[T Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func abs[T Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
