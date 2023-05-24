package gollision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBitmap_Good(t *testing.T) {
	testCases := []struct {
		desc           string
		h, w           int
		data           [][]uint8
		expectedLen    int
		expectedValues []uint64
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
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			bm := newBitmap(tc.h, tc.w, tc.data)
			assert.Equal(t, tc.expectedLen, len(bm.m), "mismatch map length")
			for i, expected := range tc.expectedValues {
				assert.Equal(t, expected, bm.m[i], "mismatch map value at %d, expected: %d, actual %d", i, expected, bm.m[i])
			}
		})
	}
}

func TestAnd_Good(t *testing.T) {
	testCases := []struct {
		desc           string
		h, w           int
		data1          [][]uint8
		data2          [][]uint8
		expectedValues []uint64
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
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			bm1 := newBitmap(tc.h, tc.w, tc.data1)
			bm2 := newBitmap(tc.h, tc.w, tc.data2)
			bm := bm1.and(bm2)
			for i, expected := range tc.expectedValues {
				assert.Equal(t, expected, bm.m[i], "mismatch map value at %d, expected: %d, actual %d", i, expected, bm.m[i])
			}
		})
	}
}

func TestOr_Good(t *testing.T) {
	testCases := []struct {
		desc           string
		h, w           int
		data1          [][]uint8
		data2          [][]uint8
		expectedValues []uint64
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
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			bm1 := newBitmap(tc.h, tc.w, tc.data1)
			bm2 := newBitmap(tc.h, tc.w, tc.data2)
			bm := bm1.or(bm2)
			for i, expected := range tc.expectedValues {
				assert.Equal(t, expected, bm.m[i], "mismatch map value at %d, expected: %d, actual %d", i, expected, bm.m[i])
			}
		})
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
		desc           string
		x, y           int
		data           [][]uint8
		expectedValues []uint64
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
			[]uint64{4},
		},
		{
			"move top to empty",
			0, -3, data,
			[]uint64{},
		},
		{
			"move bottom",
			0, 2, data,
			[]uint64{0, 0, 1, 2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			bm := newBitmap(h, w, tc.data)
			moved := bm.offset(tc.x, tc.y)
			for i, expected := range tc.expectedValues {
				assert.Equal(t, max(0, h+tc.y), len(moved.m))
				assert.Equal(t, expected, moved.m[i], "mismatch map value at %d, expected: %d, actual %d", i, expected, moved.m[i])
			}
		})
	}
}

func TestIsEmpty_Good(t *testing.T) {
	testCases := []struct {
		desc    string
		h, w    int
		data    [][]uint8
		isEmpty bool
	}{
		{
			"empty 1",
			5, 5,
			[][]uint8{},
			true,
		},
		{
			"empty 2",
			0, 0,
			[][]uint8{},
			true,
		},
		{
			"empty 3",
			0, 0,
			[][]uint8{{1, 2, 3}},
			true,
		},
		{
			"empty 4",
			5, 0,
			[][]uint8{{1, 2, 3}},
			true,
		},
		{
			"empty 5",
			0, 5,
			[][]uint8{{1, 2, 3}},
			true,
		},
		{
			"not empty 1",
			2, 3,
			[][]uint8{{1, 1, 1}, {1, 1, 1}},
			false,
		},
		{
			"not empty 2",
			2, 2,
			[][]uint8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			bm := newBitmap(tc.h, tc.w, tc.data)
			assert.Equal(t, tc.isEmpty, bm.isEmpty())
		})
	}
}
