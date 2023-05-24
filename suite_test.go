package gollision

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GollisionSuite struct {
	suite.Suite
}

func TestGollisionSuite(t *testing.T) {
	suite.Run(t, new(GollisionSuite))
}

type image struct {
	data [][]uint8
}

func (i image) w() int {
	if len(i.data) == 0 {
		return 0
	}
	return len(i.data[0])
}

func (i image) h() int {
	return len(i.data)
}

var (
	img1 = image{
		data: [][]uint8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 1, 0},
		},
	}

	img2 = image{
		data: [][]uint8{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
	}
	img3 = image{
		data: [][]uint8{
			{1, 0, 0, 0, 1},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0},
			{1, 0, 0, 0, 1},
		},
	}
)

func (su *GollisionSuite) TestIntegration() {
	sp := NewSpace()

	player := Type(0)
	monster := Type(1)

	body1 := NewBody(sp, player)
	body1.UpdateBitmap(img1.h(), img1.w(), img1.data)
	body2 := NewBody(sp, monster)
	body2.UpdateBitmap(img2.h(), img2.w(), img2.data)
	body3 := NewBody(sp, monster)
	body3.UpdateBitmap(img3.h(), img3.w(), img3.data)

	testCases := []struct {
		desc                string
		body1Dx, body1Dy    int
		expectCollidedCount [3]int
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
		su.T().Run(tc.desc, func(t *testing.T) {
			x, y := body1.UpdatePosition(tc.body1Dx, tc.body1Dy)
			su.T().Logf("%s (%d, %d)", tc.desc, x, y)
			sp.Update()
			su.Equal(tc.expectCollidedCount[0], len(body1.GetCollided()), "img%d", 1)
			su.Equal(tc.expectCollidedCount[1], len(body2.GetCollided()), "img%d", 2)
			su.Equal(tc.expectCollidedCount[2], len(body3.GetCollided()), "img%d", 3)
		})
	}
}

func TestToSlice(t *testing.T) {
	bm := newBitmap(5, 3, [][]uint8{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	})

	sli := bm.toSlice()
	expectedValues := []uint8{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1}
	for i, v := range sli {
		assert.Equal(t, expectedValues[i], v)
	}
}
