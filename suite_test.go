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
	Data [][]uint8
}

var (
	img1 = Image{
		W: 3,
		H: 3,
		Data: [][]uint8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 1, 0},
		},
	}

	img2 = Image{
		W: 5,
		H: 5,
		Data: [][]uint8{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
	}
	img3 = Image{
		W: 5,
		H: 5,
		Data: [][]uint8{
			{1, 0, 0, 0, 1},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0},
			{1, 0, 0, 0, 1},
		},
	}
)

func (su *GollisionSuite) TestIntegration() {
	return
	sp := NewSpace()

	player := Type(0)
	monster := Type(1)

	body1 := NewBody(sp, player)
	su.Require().NoError(body1.UpdateBitmap(img1.H, img1.W, img1.Data))
	body2 := NewBody(sp, monster)
	su.Require().NoError(body2.UpdateBitmap(img2.H, img2.W, img2.Data))
	body3 := NewBody(sp, monster)
	su.Require().NoError(body3.UpdateBitmap(img3.H, img3.W, img3.Data))

	testCases := []struct {
		Name                   string
		Body1MoveX, Body1MoveY int
		ExpectCollidedCount    [3]int
	}{
		{
			"Start",
			0, 0,
			[3]int{1, 1, 0},
		},
		{
			"move right",
			1, 0,
			[3]int{1, 0, 1},
		},
		{
			"move left",
			-1, 0,
			[3]int{1, 1, 0},
		},
		{
			"move up",
			0, -1,
			[3]int{2, 1, 1},
		},
		{
			"move bottom",
			0, 1,
			[3]int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		_, _ = body1.UpdatePosition(tc.Body1MoveX, tc.Body1MoveY)
		sp.Update()
		su.Equal(tc.ExpectCollidedCount[0], len(body1.GetCollided()), "%s: %d", tc.Name, 1)
		su.Equal(tc.ExpectCollidedCount[1], len(body2.GetCollided()), "%s: %d", tc.Name, 2)
		su.Equal(tc.ExpectCollidedCount[2], len(body3.GetCollided()), "%s: %d", tc.Name, 3)
	}
}
