package gollision

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GollisionSuite struct {
	suite.Suite
}

func TestGollisionSuite(t *testing.T) {
	suite.Run(t, new(GollisionSuite))
}

type Image struct {
	W, H int
	Data []uint8
}

func (i Image) Vectors() []Vector {
	var vectors []Vector
	for y := 0; y < i.H; y++ {
		for x := 0; x < i.W; x++ {
			if i.Data[y*i.W+x] != 0 {
				vectors = append(vectors, Vector{X: x, Y: y})
			}
		}
	}
	return vectors
}

var (
	img1 = Image{
		W: 3,
		H: 3,
		Data: []uint8{
			0, 0, 0,
			0, 0, 0,
			0, 1, 0,
		},
	}

	img2 = Image{
		W: 5,
		H: 5,
		Data: []uint8{
			0, 0, 0, 0, 0,
			0, 1, 1, 1, 0,
			0, 1, 0, 1, 0,
			0, 1, 1, 1, 0,
			0, 0, 0, 0, 0,
		},
	}
	img3 = Image{
		W: 5,
		H: 5,
		Data: []uint8{
			1, 0, 0, 0, 1,
			0, 1, 0, 1, 0,
			0, 0, 1, 0, 0,
			0, 1, 0, 1, 0,
			1, 0, 0, 0, 1,
		},
	}
)

func (s *GollisionSuite) TestIntegration() {
	sp := NewSpace()

	player := Type(0)
	monster := Type(1)

	body1 := NewBody(&sp, NewBitmap2D(img1.W, img1.H, img1.Vectors()...), player, 0, 0)
	body2 := NewBody(&sp, NewBitmap2D(img2.W, img2.H, img2.Vectors()...), monster, 0, 0)
	body3 := NewBody(&sp, NewBitmap2D(img3.W, img3.H, img3.Vectors()...), monster, 0, 0)

	sp.Update()

	debug("Body1 Hit", body1.GetCollided())
	debug("Body2 Hit", body2.GetCollided())
	debug("Body3 Hit", body3.GetCollided())

	println()
	body1.UpdatePosition(Vector{X: 1})
	sp.Update()

	debug("Body1 Hit", body1.GetCollided())
	debug("Body2 Hit", body2.GetCollided())
	debug("Body3 Hit", body3.GetCollided())
}

func debug(prefix string, obj []Body) {
	print(prefix)
	for _, o := range obj {
		print(" ")
		print(o.ID())
	}
	if len(obj) == 0 {
		print(" ")
		print("nothing")
	}
	println("")
}
