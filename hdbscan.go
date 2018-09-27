package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func loadData(filePath string) [][]float32 {
	arr := [][]float32{
		[]float32{0.0, 0.0},
		[]float32{0.0, 0.0},
	}
	return arr
}

func saveClusteringResult(savePath string, data [][]float32) {
	file, _ := os.Create(savePath)
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()
	for _, element := range data {
		st := strings.Fields(strings.Trim(fmt.Sprint(element), "[]"))
		wr.Write(st)
	}
}

func cluster(dataPath string, savePath string, minClusterSize int, metric string, alpha float64, algorithm string, leafSize int, genMinSpanTree bool, clusterSelectionMethod string) {
	fmt.Println("Loading data")
	data := loadData(dataPath)
	fmt.Println("Saving data")
	saveClusteringResult("blah.csv", data)
}
