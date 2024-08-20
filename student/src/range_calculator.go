package src

import (
	"math"
	"slices"
)

// Average takes a slice of float64 and returns the average as a float64.
func Average(nums []float64) float64 {
	sum := 0.0
	for _, value := range nums {
		sum += value
	}
	return sum / float64(len(nums))
}

// Median takes a slice of float64 and returns the median as a float64.
func Median(numbers []float64) float64 {
	slices.Sort(numbers)
	leng := len(numbers)
	if leng%2 != 0 {
		return numbers[leng/2-1]
	} else {
		mid1 := numbers[(leng/2)-1]
		mid2 := numbers[leng/2]
		return (mid1 + mid2) / 2
	}
}

// Variance calculates the variance of a slice of float64.
func Variance(nums []float64) float64 {
	sumSquared := 0.0
	mean := Average(nums)
	for _, value := range nums {
		sumSquared += (value - mean) * (value - mean)
	}
	return sumSquared / float64(len(nums))
	
}

// StdDev calculates the standard deviation of a slice of float64.
func StdDev(nums []float64) float64 {
	return math.Sqrt(Variance(nums))
}

// CalculateRange calculates the predicted range for the next number.
func CalculateRange(nums []float64, isFirst bool) (float64, float64) {
	if isFirst {
		// Provide a wider range for the first number
		return -10, 10 // Adjust these values as needed for wider initial range
	}

	mean := Average(nums)
	stdDev := StdDev(nums)

	// Use a fraction of the standard deviation to tighten the range
	rangeFactor := 0.99 // Adjust this factor as needed to tighten the range
	lower := mean - (rangeFactor * stdDev*2)
	upper := mean + (rangeFactor * stdDev*2)

	// Ensure the lower bound doesn't go below zero
	if lower < 0 {
		lower = 0
	}

	return lower, upper
}
