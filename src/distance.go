package main

import (
	"math"
)

func euclidianDistance(p1 []float64, p2 []float64) float64 {
	acc := 0.0
	for i, v := range p1 {
		acc += math.Pow((v - p2[i]), 2)
	}
	return math.Pow(acc, 0.5)
}
