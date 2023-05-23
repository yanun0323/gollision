package gollision

type Body interface {
	ID() uint64
	Type() Type
	SetPosition(x, y int)
	UpdatePosition(dx, dy int) (x, y int)
	UpdateBitmap(h, w int, image [][]uint8) error
	GetCollided() []Body

	positionedBitmap() bitmap
}

type Type uint64

type body struct {
	id   uint64
	s    Space
	t    Type
	bm   bitmap
	x, y int
}

func NewBody(s Space, t Type) Body {
	b := &body{
		id: s.nextID(),
		s:  s,
		t:  t,
	}

	s.addBody(b)
	return b
}

func (b body) ID() uint64 {
	return b.id
}

func (b body) Type() Type {
	return b.t
}
func (b *body) SetPosition(x, y int) {
	b.x = x
	b.y = y
}
func (b *body) UpdatePosition(dx, dy int) (x, y int) {
	b.x = x + dx
	b.y = y + dy
	return b.x, b.y
}

func (b *body) UpdateBitmap(h, w int, image [][]uint8) error {
	bm, err := newBitmap(h, w, image)
	if err != nil {
		return err
	}
	b.bm = bm
	return nil
}

func (b *body) GetCollided() []Body {
	return b.s.GetCollided(b.id)
}

func (b *body) positionedBitmap() bitmap {
	if b.bm.isEmpty() {
		return bitmap{}
	}
	return b.bm.offset(b.x, b.y)
}
