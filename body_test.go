package gollision

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBody(t *testing.T) {
	s := NewSpace()
	for i := 0; i < 100; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NewBody(s, 0)
			assert.Equal(t, uint64(i+1), b.ID())
		})
	}
}

func TestID(t *testing.T) {
	var b Body = &body{id: 10}
	assert.Equal(t, uint64(10), b.ID())
}

func TestType(t *testing.T) {
	b := &body{t: 30}
	var bb Body = b
	assert.Equal(t, Type(30), bb.Type())
}

func TestSetPosition(t *testing.T) {
	b := &body{x: 10, y: 10}
	var bb Body = b
	bb.SetPosition(13, 14)
	assert.Equal(t, 13, b.x)
	assert.Equal(t, 14, b.y)
}

func TestAddPosition(t *testing.T) {
	b := &body{x: 10, y: 10}
	var bb Body = b
	x, y := bb.AddPosition(13, 14)
	assert.Equal(t, 23, x)
	assert.Equal(t, 23, b.x)
	assert.Equal(t, 24, y)
	assert.Equal(t, 24, b.y)
}

func TestUpdateBitmapByImage(t *testing.T) {
	// TODO: Implement me
}

func TestUpdateBitmapByAlpha(t *testing.T) {
	// TODO: Implement me
}

func TestIsCollided(t *testing.T) {
	// TODO: Implement me
}

func TestGetCollided(t *testing.T) {
	// TODO: Implement me
}

func TestPositionedBitmap(t *testing.T) {
	// TODO: Implement me
}
