package main

import (
	"fmt"
)

func max(arr []float64) float64 {
	max := 0.0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func printTreeStructure(tree *node) {
	fmt.Println(tree)
	for _, c := range tree.children {
		printTreeStructure(c)
	}
}

func printClusterStructure(tree *cluster) {
	fmt.Println(tree)
	for _, c := range tree.children {
		printClusterStructure(c)
	}
}
