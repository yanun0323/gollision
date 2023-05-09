package gollision

type Body interface {
	ID() uint64
	Type() Type

	Bitmap() Bitmap2D
	PositionedBitmap() Bitmap2D

	SetPosition(v Vector)
	UpdatePosition(offset Vector) Vector

	UpdateBitmap(bm Bitmap2D)

	GetCollided() []Body
}
