package main

import (
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	maxCount := 0
	for i := 0; i < len(points); i++ {
		// A map gradientCounts is used to keep track of the count of each gradient.
		// If the count of a particular gradient in gradientCounts is greater than
		// the current maxInGroup, maxInGroup is updated.
		maxInGroup := 0
		gradientCounts := make(map[float64]int)
		for j := i + 1; j < len(points); j++ {
			gradient := calculateGradient(points[i], points[j])
			// corrects golang +/- zero
			gradient = normalizeZeroAndInf(gradient)

			gradientCounts[gradient]++

			if maxInGroup < gradientCounts[gradient] {
				maxInGroup = gradientCounts[gradient]
			}
		}
		maxCount = max(maxCount, maxInGroup)

		// optimize the algorithm. It's a form of early stopping.
		if maxCount+1 == len(points)-i {
			return maxCount + 1
		}
	}

	//  maxCount + 1 is returned because we are grouping the points by line, not by points.
	// The 1 accounts for the point itself.
	return maxCount + 1
}

// max is a function that returns the maximum of two integers.
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// calculateGradient is a function that calculates the gradient between two points.
// It returns an error if the difference in x-coordinates is zero to prevent division by zero.
func calculateGradient(point1, point2 []int) float64 {
	result := float64(0)
	result = float64(point1[1]-point2[1]) / float64(point1[0]-point2[0])
	return result
}

// normalizeZeroAndInf is a function that takes a float64 number and if it's zero or infinity
// (positive or negative), it returns its absolute value.
// This is done because Go distinguishes between positive and negative zeros and infinities,
// but for the purpose of this function, we consider them equivalent and thus normalize them
// to their positive versions.
func normalizeZeroAndInf(gradient float64) float64 {
	if gradient == 0 || math.IsInf(gradient, 0) {
		return math.Abs(gradient)
	}
	return gradient
}
