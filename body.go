package gollision

type Type uint64

type Vector struct {
	X, Y int
}

func (v Vector) Add(offset Vector) Vector {
	v.X += offset.X
	v.Y += offset.Y
	return v
}

type body struct {
	id uint64
	s  *space
	t  Type
	bm Bitmap2D
	v  Vector
}

func NewBody(s *space, bm Bitmap2D, t Type, x, y int) Body {
	b := &body{
		id: s.NextID(),
		s:  s,
		t:  t,
		bm: bm,
		v:  Vector{X: x, Y: y},
	}
	s.AddBody(b)
	return b
}

func (b body) ID() uint64 {
	return b.id
}

func (b body) Type() Type {
	return b.t
}

func (b *body) Bitmap() Bitmap2D {
	return b.bm
}

func (b *body) PositionedBitmap() Bitmap2D {
	return b.bm.Offset(b.v.X, b.v.Y)
}
func (b *body) SetPosition(v Vector) {
	b.v = v
}
func (b *body) UpdatePosition(offset Vector) Vector {
	b.v = b.v.Add(offset)
	return b.v
}

func (b *body) UpdateBitmap(bm Bitmap2D) {
	b.bm = bm
}

func (b *body) GetCollided() []Body {
	return b.s.GetCollided(b.id)
}
