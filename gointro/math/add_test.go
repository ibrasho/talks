package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		x, y     int
		expected int
	}{
		{0, 0, 0},
		{1, 2, 3},
		{5, 5, 10},
		{5, -5, 0},
		{-5, 5, 0},
		{-5, -5, -10},
	}

	for _, c := range cases {
		if actual := Add(c.x, c.y); actual != c.expected {
			t.Errorf("Add(%v, %v) faield: expected: %v, got: %v",
				c.x, c.y, c.expected, actual)
		}
	}
}
