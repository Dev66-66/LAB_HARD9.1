package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{2, 4},
		{-3, 9},
		{10, 100},
	}

	for _, tt := range tests {
		got := Square(tt.n)
		if got != tt.want {
			t.Errorf("Square(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
