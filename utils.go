package main

func max(arr []float64) float64 {
	max := 0.0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
