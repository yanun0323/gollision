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

func (su *GollisionSuite) TestIntegration() {
	sp := NewSpace()

	player := Type(0)
	monster := Type(1)

	body1 := NewBody(&sp, NewBitmap2D(img1.W, img1.H, img1.Vectors()...), player, 0, 0)
	body2 := NewBody(&sp, NewBitmap2D(img2.W, img2.H, img2.Vectors()...), monster, 0, 0)
	body3 := NewBody(&sp, NewBitmap2D(img3.W, img3.H, img3.Vectors()...), monster, 0, 0)

	testCases := []struct {
		Name                string
		Body1Move           Vector
		ExpectCollidedCount [3]int
	}{
		{
			"Start",
			Vector{},
			[3]int{1, 1, 0},
		},
		{
			"move right",
			Vector{X: 1},
			[3]int{1, 0, 1},
		},
		{
			"move left",
			Vector{X: -1},
			[3]int{1, 1, 0},
		},
		{
			"move up",
			Vector{Y: -1},
			[3]int{2, 1, 1},
		},
		{
			"move bottom",
			Vector{Y: 1},
			[3]int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		body1.UpdatePosition(tc.Body1Move)
		sp.Update()
		su.Equal(tc.ExpectCollidedCount[0], len(body1.GetCollided()), "%s: %d", tc.Name, 1)
		su.Equal(tc.ExpectCollidedCount[1], len(body2.GetCollided()), "%s: %d", tc.Name, 2)
		su.Equal(tc.ExpectCollidedCount[2], len(body3.GetCollided()), "%s: %d", tc.Name, 3)
	}
}
