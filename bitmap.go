package gollision

type Bitmap2D interface {
	Width() int
	Height() int
	Bitmap() []uint64
	And(in Bitmap2D) Bitmap2D
	IsZero() bool
	Offset(x int, y int) Bitmap2D
	Or(in Bitmap2D) Bitmap2D
	ToSlice() []uint8
}

type bitmap2D struct {
	w, h int
	m    []uint64
}

func NewBitmap2D(w, h int, vectors ...Vector) Bitmap2D {
	m := make([]uint64, h)
	for _, vec := range vectors {
		m[vec.Y] |= uint64(1) << vec.X
	}
	return &bitmap2D{
		w: w,
		h: h,
		m: m,
	}
}

func (bm *bitmap2D) Width() int {
	return bm.w
}
func (bm *bitmap2D) Height() int {
	return bm.h
}
func (bm *bitmap2D) Bitmap() []uint64 {
	return bm.m
}

func (bm *bitmap2D) And(in Bitmap2D) Bitmap2D {
	minH := min(bm.h, in.Height())
	inMap := in.Bitmap()
	m := make([]uint64, minH)
	for i := 0; i < minH; i++ {
		m[i] = bm.m[i] & inMap[i]
	}
	return &bitmap2D{
		w: min(bm.w, in.Width()),
		h: minH,
		m: m,
	}
}

func (bm *bitmap2D) Or(in Bitmap2D) Bitmap2D {
	higher := in
	if in.Height() > bm.h {
		higher = bm
	}

	m := make([]uint64, higher.Height())
	higherMap := higher.Bitmap()
	for i := range higherMap {
		if i < higher.Height() {
			m[i] = bm.m[i] | higherMap[i]
			continue
		}
		m[i] = higherMap[i]
	}
	return &bitmap2D{
		w: max(bm.w, in.Width()),
		h: higher.Height(),
		m: m,
	}
}

func (bm *bitmap2D) Offset(x, y int) Bitmap2D {
	w := bm.w + x
	h := bm.h + y
	if w <= 0 || h <= 0 {
		return NewBitmap2D(0, 0)
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
	return &bitmap2D{
		w: w,
		h: h,
		m: m,
	}
}

func (bm *bitmap2D) IsZero() bool {
	for _, v := range bm.m {
		if v != 0 {
			return false
		}
	}
	return true
}

func (bm *bitmap2D) ToSlice() []uint8 {
	s := make([]uint8, bm.w*bm.h)
	for i := range s {
		if bm.m[i/bm.w]&(uint64(1)<<uint(i%bm.w)) != 0 {
			s[i] = 1
		}
	}
	return s
}
