package main

import (
	"fmt"
	"math"
)

var tInitial float64 = 0
var tFinal float64 = 10

var yInitial float64 = 1

func evalYExpected(t float64) float64 {
	return math.Cos(t)
}

func evalF(t float64) float64 {
	return math.Sin(t)
}

/**
 * Solves the differential equation and returns the approximate solution as well
 * as the error.
 */
func solve(N int) ([]float64, float64) {
	var error float64 = 0

	deltaT := (tFinal - tInitial) / float64(N)

	yArray := make([]float64, N+1)
	yArray[0] = yInitial

	for i := 1; i <= N; i++ {
		t := tInitial + deltaT*float64(i)

		yArray[i] = yArray[i-1] + evalF(t)*deltaT

		thisError := math.Abs(yArray[i] - evalYExpected(t))

		if thisError > error {
			error = thisError
		}
	}

	return yArray, error
}

var minN = 16
var maxN = 2048

func main() {
	var prevError float64 = 0

	for N := minN; N <= maxN; N *= 2 {
		_, error := solve(N)

		if N != minN {
			convergence := math.Abs(error / prevError)

			fmt.Printf("N = %d: %g\n", N, convergence)
		}

		prevError = error
	}
}
