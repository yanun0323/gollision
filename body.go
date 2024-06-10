package gollision

type Body interface {
	// The ID of body
	ID() uint64

	// The type of body
	Type() Type

	// It will change position to (x, y).
	SetPosition(x, y int)

	// It will change position to (origin x + dy, origin y + dy),
	// and return position that body moved to.
	AddPosition(dx, dy int) (x, y int)

	// Boundary gets the size of collision body
	Boundary() (w, h int)

	// Update body image data
	UpdateBitmapByImage(h, w int, image [][]uint8)

	// Update body alpha data
	UpdateBitmapByAlpha(h, w int, image [][]bool)

	// Return true if this body is colliding something
	IsCollided() bool

	// Return the other bodies colliding with this body
	GetCollided() []Body

	// Return the bitmap witch added the position offset of this body
	positionedBitmap() *bitmap
}

type Type uint8

type body struct {
	id   uint64
	s    Space
	t    Type
	bm   *bitmap
	x, y int
}

// Create a new body and add it into the space
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

func (b body) Boundary() (w, h int) {
	return b.bm.w, b.bm.h
}

func (b *body) SetPosition(x, y int) {
	b.x = x
	b.y = y
}

func (b *body) AddPosition(dx, dy int) (x, y int) {
	b.x += dx
	b.y += dy
	return b.x, b.y
}

func (b *body) UpdateBitmapByImage(h, w int, image [][]uint8) {
	bm := newBitmapByImage(h, w, image)
	b.bm = bm
}

func (b *body) UpdateBitmapByAlpha(h, w int, image [][]bool) {
	bm := newBitmapByAlpha(h, w, image)
	b.bm = bm
}

func (b *body) IsCollided() bool {
	return b.s.IsCollided(b.id)
}

func (b *body) GetCollided() []Body {
	return b.s.GetCollided(b.id)
}

func (b *body) positionedBitmap() *bitmap {
	if b.bm.isEmpty() {
		return emptyBitmap()
	}
	return b.bm.offset(b.x, b.y)
}
