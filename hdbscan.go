package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadData(filePath string) [][]float64 {
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	arr := [][]float64{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		subArr := []float64{}
		for _, element := range line {
			fl, _ := strconv.ParseFloat(element, 64)
			subArr = append(subArr, fl)
		}
		arr = append(arr, subArr)
	}
	return arr
}

func saveClusteringResult(savePath string, data [][]float64) {
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

	// Find all clusters of size minClusterSize

	fmt.Println("Saving data")
	saveClusteringResult(savePath, data)
}
