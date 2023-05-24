package gollision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax_Good(t *testing.T) {
	testCases := []struct {
		desc           string
		a, b, expected int
	}{
		{"test A", 2, 1, 2},
		{"test B", -2, 1, 1},
		{"test C", -2, -1, -1},
		{"test D", 2, -1, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			assert.Equal(t, tc.expected, max(tc.a, tc.b))
		})
	}
}

func TestMin_Good(t *testing.T) {
	testCases := []struct {
		desc           string
		a, b, expected int
	}{
		{"test A", 2, 1, 1},
		{"test B", -2, 1, -2},
		{"test C", -2, -1, -2},
		{"test D", 2, -1, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			assert.Equal(t, tc.expected, min(tc.a, tc.b))
		})
	}
}

func TestAbs_Good(t *testing.T) {
	testCases := []struct {
		desc        string
		x, expected int
	}{
		{"test A", 2, 2},
		{"test B", -2, 2},
		{"test C", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Log(tc.desc)
			assert.Equal(t, tc.expected, abs(tc.x))
		})
	}
}
