package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"allan/src"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data []float64
	windowSize := 10 // Increase the window size for better accuracy

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue // skip non-numeric values
		}
		data = append(data, value)

		// Ensure that we're only using the last 'windowSize' numbers
		if len(data) > windowSize {
			data = data[len(data)-windowSize:]
		}

		// Only calculate if we have enough data points
		if len(data) > 1 {
			mean := src.Mean(data)
			stDev := math.Sqrt(src.Variance(data))

			// Use a very high multiplier for 99.99% confidence
			lowerBound := int(mean - 3.89*stDev) // Slightly more than 3.89
			upperBound := int(mean + 3.89*stDev)

			fmt.Printf("%d %d\n", lowerBound, upperBound)
		}
	}
}
