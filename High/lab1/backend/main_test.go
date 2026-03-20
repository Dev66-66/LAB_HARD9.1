package main

import (
	"testing"
)

func TestHeavyComputation(t *testing.T) {
	tests := []struct {
		name       string
		start      int64
		end        int64
		workers    int
		wantResult int64
	}{
		{"1 to 3, 1 worker", 1, 3, 1, 1*1 + 2*2 + 3*3}, // 1+4+9=14
		{"1 to 10, 4 workers", 1, 10, 4, 385},
		{"Single number", 5, 5, 1, 25},
		{"Empty range", 10, 5, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HeavyComputation(tt.start, tt.end, tt.workers)
			if got != tt.wantResult {
				t.Errorf("HeavyComputation() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
