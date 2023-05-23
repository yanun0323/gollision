package gollision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBitmap_Good(t *testing.T) {
	testCases := []struct {
		Name           string
		H, W           int
		Data           [][]uint8
		ExpectedLen    int
		ExpectedValues []uint64
	}{
		{
			"empty 1", 0, 0, [][]uint8{},
			0, []uint64{},
		},
		{
			"empty 2", 0, 0, [][]uint8{{1, 1, 1}},
			0, []uint64{},
		},
		{
			"empty 3", 3, 0, [][]uint8{},
			0, []uint64{},
		},
		{
			"empty 4", 3, 0, [][]uint8{{1, 1, 1}},
			0, []uint64{},
		},
		{
			"empty 5", 0, 3, [][]uint8{},
			0, []uint64{},
		},
		{
			"empty 6", 0, 3, [][]uint8{{1, 1, 1}},
			0, []uint64{},
		},
		{
			"general", 5, 3, [][]uint8{
				{0, 0, 0, 255},
				{0, 0, 1, 0},
				{0, 0, 255, 255},
				{255, 0, 0, 0},
				{255, 255, 0, 0},
				{1, 1, 1, 255},
			},
			5,
			[]uint64{0, 4, 4, 1, 3},
		},
	}

	for _, tc := range testCases {
		bm, err := newBitmap(tc.H, tc.W, tc.Data)
		assert.NoError(t, err, "test: %s, err: %+v", tc.Name, err)
		assert.Equal(t, tc.ExpectedLen, len(bm.m), "test: %s, mismatch map length", tc.Name)
		for i, expected := range tc.ExpectedValues {
			assert.Equal(t, expected, bm.m[i], "test: %s, mismatch map value at %d, expected: %d, actual %d", tc.Name, i, expected, bm.m[i])
		}
	}
}

func TestNewBitmap_Error(t *testing.T) {
	testCases := []struct {
		Name string
		H, W int
		Data [][]uint8
	}{
		{
			"height out of bounds", 5, 3, [][]uint8{
				{0, 0, 0},
				{0, 0, 1},
				{0, 0, 255},
			},
		},
		{
			"width out of bounds", 3, 5, [][]uint8{
				{0, 0, 0},
				{0, 0, 1},
				{0, 0, 255},
			},
		},
		{
			"empty 1", 3, 3, [][]uint8{},
		},
	}

	for _, tc := range testCases {
		_, err := newBitmap(tc.H, tc.W, tc.Data)
		assert.Error(t, err, "test: %s no error", tc.Name)
	}
}

func TestAnd_Good(t *testing.T) {
	testCases := []struct {
		Name     string
		H, W     int
		Data1    [][]uint8
		Data2    [][]uint8
		Expected []uint64
	}{
		{
			"full test",
			5, 3,
			[][]uint8{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			[][]uint8{
				{0, 0, 0, 111},
				{255, 0, 0, 111},
				{0, 0, 255, 111},
				{255, 255, 0, 111},
				{255, 255, 255, 111},
			},
			[]uint64{0, 1, 4, 3, 7},
		},
	}

	for _, tc := range testCases {
		bm1, err := newBitmap(tc.H, tc.W, tc.Data1)
		assert.NoError(t, err, tc.Name)
		bm2, err := newBitmap(tc.H, tc.W, tc.Data2)
		assert.NoError(t, err, tc.Name)

		bm := bm1.and(bm2)
		for i, expected := range tc.Expected {
			assert.Equal(t, expected, bm.m[i], "test: %s, mismatch map value at %d, expected: %d, actual %d", tc.Name, i, expected, bm.m[i])

		}
	}
}

func TestOr_Good(t *testing.T) {
	testCases := []struct {
		Name     string
		H, W     int
		Data1    [][]uint8
		Data2    [][]uint8
		Expected []uint64
	}{
		{
			"full test",
			5, 3,
			[][]uint8{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			[][]uint8{
				{0, 0, 0, 111},
				{255, 0, 0, 111},
				{0, 0, 255, 111},
				{255, 255, 0, 111},
				{255, 255, 255, 111},
			},
			[]uint64{0, 1, 4, 3, 7},
		},
	}

	for _, tc := range testCases {
		bm1, err := newBitmap(tc.H, tc.W, tc.Data1)
		assert.NoError(t, err, tc.Name)
		bm2, err := newBitmap(tc.H, tc.W, tc.Data2)
		assert.NoError(t, err, tc.Name)

		bm := bm1.or(bm2)
		for i, expected := range tc.Expected {
			assert.Equal(t, expected, bm.m[i], "test: %s, mismatch map value at %d, expected: %d, actual %d", tc.Name, i, expected, bm.m[i])
		}
	}
}

func TestOffset_Good(t *testing.T) {
	data := [][]uint8{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	h, w := len(data), len(data[0])

	testCases := []struct {
		Name           string
		X, Y           int
		Data           [][]uint8
		ExpectedValues []uint64
	}{
		{
			"move right",
			2, 0, data,
			[]uint64{4, 8, 16},
		},
		{
			"move left",
			-2, 0, data,
			[]uint64{0, 0, 1},
		},
		{
			"move left to empty",
			-3, 0, data,
			[]uint64{},
		},
		{
			"move top",
			0, -2, data,
			[]uint64{16},
		},
		{
			"move top to empty",
			0, -3, data,
			[]uint64{},
		},
		{
			"move bottom",
			0, 2, data,
			[]uint64{0, 0, 4, 8, 16},
		},
	}

	for _, tc := range testCases {
		bm, err := newBitmap(h, w, tc.Data)
		assert.NoError(t, err, tc.Name)
		moved := bm.offset(tc.X, tc.Y)
		for i, expected := range tc.ExpectedValues {
			assert.Equal(t, max(0, h+tc.Y), len(moved.m), tc.Name)
			assert.Equal(t, expected, moved.m[i], "test: %s, mismatch map value at %d, expected: %d, actual %d", tc.Name, i, expected, moved.m[i])
		}
	}
}
