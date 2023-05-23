package gollision

import "errors"

type bitmap struct {
	w, h  int
	m     []uint64
	empty bool
}

func newBitmap(h, w int, data [][]uint8) (bitmap, error) {
	if h == 0 || w == 0 {
		return bitmap{}, nil
	}

	if len(data) < h {
		return bitmap{}, errors.New("height out of bounds")
	}

	if len(data) == 0 || len(data[0]) < w {
		return bitmap{}, errors.New("width out of bounds")
	}

	m := make([]uint64, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if data[y][x] == 0 {
				continue
			}
			m[y] |= uint64(1) << x
		}
	}

	return bitmap{
		w: w,
		h: h,
		m: m,
	}, nil
}

func (bm bitmap) and(in bitmap) bitmap {
	minH := min(bm.h, in.h)
	m := make([]uint64, minH)
	for i := 0; i < minH; i++ {
		m[i] = bm.m[i] & in.m[i]
	}
	return bitmap{
		w: min(bm.w, in.w),
		h: minH,
		m: m,
	}
}

func (bm bitmap) or(in bitmap) bitmap {
	higher := in
	if in.h > bm.h {
		higher = bm
	}
	minH := min(bm.h, in.h)

	m := make([]uint64, higher.h)
	higherMap := higher.m
	for i := range higherMap {
		if i < minH {
			m[i] = bm.m[i] | in.m[i]
			continue
		}
		m[i] = higherMap[i]
	}
	return bitmap{
		w: max(bm.w, in.w),
		h: higher.h,
		m: m,
	}
}

func (bm bitmap) offset(x, y int) bitmap {
	w := bm.w + x
	h := bm.h + y
	if w <= 0 || h <= 0 {
		return bitmap{}
	}

	temp := bm.m
	m := make([]uint64, h)
	if y < 0 {
		temp = temp[abs(y):]
		y = 0
	}

	for i := range temp {
		if x >= 0 {
			m[i+y] = temp[i] << x
		} else {
			m[i+y] = temp[i] >> abs(x)
		}
	}
	return bitmap{
		w: w,
		h: h,
		m: m,
	}
}

func (bm *bitmap) isEmpty() bool {
	if bm.empty {
		return true
	}

	for _, v := range bm.m {
		if v != 0 {
			return false
		}
	}

	bm.empty = true
	return true
}

func (bm bitmap) toSlice() []uint8 {
	s := make([]uint8, bm.w*bm.h)
	for i := range s {
		if bm.m[i/bm.w]&(uint64(1)<<uint(i%bm.w)) != 0 {
			s[i] = 1
		}
	}
	return s
}
