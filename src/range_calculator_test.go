package src

import "testing"

func TestMean(t *testing.T) {
	tests := []struct {
		data []float64
		expected float64
	}{
		
        {data: []float64{1, 2, 3, 4, 5}, expected: 3.0},
        {data: []float64{10, 20, 30}, expected: 20.0},
        {data: []float64{-1, -2, -3}, expected: -2.0},// TODO: Add test cases.
	}
	for _, tt := range tests {
        result := Mean(tt.data)
        if result != tt.expected {
            t.Errorf("Mean(%v) = %v; want %v", tt.data, result, tt.expected)
        }
    }
}

func TestVariance(t *testing.T) {
    tests := []struct {
        data     []float64
        expected float64
    }{
        {data: []float64{1, 2, 3, 4, 5}, expected: 2.0},
        {data: []float64{10, 20, 30}, expected: 66.66666666666667},
        {data: []float64{-1, -2, -3}, expected : 2.0 / 3.0},
    }

    for _, tt := range tests {
        result := Variance(tt.data)
        if result != tt.expected {
            t.Errorf("Variance(%v) = %v; want %v", tt.data, result, tt.expected)
        }
    }
}