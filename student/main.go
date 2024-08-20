package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"allan/src"
)

func main() {
	var numbers []float64
	isFirstInput := true

	for {
		var input string

		// Get number from buffer
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		// Trim any whitespace and newline characters
		input = strings.TrimSpace(input)

		// Convert the input to a float
		floatValue, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error converting input to float:", err)
			os.Exit(1)
		}

		// Append the number to the slice
		numbers = append(numbers, floatValue)

		// Provide upper and lower limit
		var lower, upper float64
		if isFirstInput {
			lower, upper = src.CalculateRange(numbers, true)
			isFirstInput = false
		} else {
			lower, upper = src.CalculateRange(numbers, false)
		}

		fmt.Printf("%.0f %.0f\n", lower, upper)
	}
}
