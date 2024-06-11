package gollision

type bitmap struct {
	w, h  int
	m     []uint64
	empty bool
}

func newBitmapByImage(h, w int, data [][]uint8) *bitmap {
	if w == 0 {
		return emptyBitmap()
	}

	m := make([]uint64, h)
	for y, row := range data {
		if y >= h {
			continue
		}
		for x, value := range row {
			if x >= w || value == 0 {
				continue
			}
			m[y] |= uint64(1) << x
		}
	}

	return &bitmap{
		w: w,
		h: h,
		m: m,
	}
}

func newBitmapByAlpha(h, w int, data [][]bool) *bitmap {
	if w == 0 {
		return emptyBitmap()
	}

	m := make([]uint64, h)
	for y, row := range data {
		if y >= h {
			continue
		}
		for x, value := range row {
			if x >= w || !value {
				continue
			}
			m[y] |= uint64(1) << x
		}
	}

	return &bitmap{
		w: w,
		h: h,
		m: m,
	}
}

func emptyBitmap() *bitmap {
	return &bitmap{m: []uint64{}, empty: true}
}

func (bm *bitmap) and(in *bitmap) *bitmap {
	if bm == nil {
		return nil
	}

	minH := min(bm.h, in.h)
	m := make([]uint64, minH)
	for i := 0; i < minH; i++ {
		m[i] = bm.m[i] & in.m[i]
	}
	return &bitmap{
		w: min(bm.w, in.w),
		h: minH,
		m: m,
	}
}

func (bm *bitmap) or(in *bitmap) *bitmap {
	if bm == nil {
		return nil
	}

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
	return &bitmap{
		w: max(bm.w, in.w),
		h: higher.h,
		m: m,
	}
}

func (bm *bitmap) offset(x, y int) *bitmap {
	if bm == nil {
		return nil
	}

	w := bm.w + x
	h := bm.h + y
	if w <= 0 || h <= 0 {
		return emptyBitmap()
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
	return &bitmap{
		w: w,
		h: h,
		m: m,
	}
}

func (bm *bitmap) isEmpty() bool {
	if bm == nil {
		return true
	}

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
