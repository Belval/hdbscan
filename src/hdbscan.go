package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type link struct {
	p1   int
	p2   int
	dist float64
}

type node struct {
	key             int
	parentKey       int
	parent          *node
	distToParent    float64
	children        []*node
	descendantCount int
}

type cluster struct {
	parent           *cluster
	points           []int
	pointsDistParent []float64
	children         []*cluster
	lBirth           float64
	lDeath           float64
	stability        float64
}

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

func computeMutualReachability(minClusterSize int, data [][]float64) [][]float64 {
	dataLen := len(data)
	fmt.Println(dataLen)
	mutualReachabilityArr := [][]float64{}
	coreDist := []float64{}
	pointDist := [][]float64{}
	for _, point := range data {
		pointDistTmp := []float64{}
		for _, subPoint := range data {
			pointDistTmp = append(pointDistTmp, euclidianDistance(point, subPoint))
		}
		pointDist = append(pointDist, pointDistTmp)
	}
	for i := 0; i < dataLen; i++ {
		pointDistTmp := []float64{}
		pointDistTmp = append(pointDistTmp, pointDist[i]...)
		sort.Float64s(pointDistTmp)
		fmt.Println(pointDistTmp)
		coreDist = append(coreDist, pointDistTmp[minClusterSize-1])
	}
	for i := 0; i < dataLen; i++ {
		mutualReachabilityArrTmp := []float64{}
		for j := 0; j < dataLen; j++ {
			mutualReachabilityArrTmp = append(mutualReachabilityArrTmp, max([]float64{coreDist[i], coreDist[j], pointDist[i][j]}))
		}
		mutualReachabilityArr = append(mutualReachabilityArr, mutualReachabilityArrTmp)
	}
	return mutualReachabilityArr
}

func getNearestPoint(linked map[int]bool, mutualReachabilityArr [][]float64) link {
	minDist := math.MaxFloat64
	p1MinIdx := 0
	p2MinIdx := 0
	for i := 0; i < len(mutualReachabilityArr); i++ {
		if _, ok := linked[i]; !ok {
			for j := range linked {
				if minDist > mutualReachabilityArr[i][j] {
					minDist = mutualReachabilityArr[i][j]
					p1MinIdx = j
					p2MinIdx = i
				}
			}
		}
	}
	return link{p1: p1MinIdx, p2: p2MinIdx, dist: mutualReachabilityArr[p1MinIdx][p2MinIdx]}
}

func computeMinSpanningTree(mutualReachabilityArr [][]float64) []link {
	linked := make(map[int]bool)
	links := []link{}
	for len(links) < len(mutualReachabilityArr) {
		link := getNearestPoint(linked, mutualReachabilityArr)
		links = append(links, link)
		linked[link.p2] = true
	}
	return links
}

func buildClusterHierarchy(links []link, leaves []node) node {
	if len(links) == 0 {
		return leaves[0]
	}
	remainingLinks := []link{}
	p1Map := make(map[int]bool)
	for _, l := range links {
		if _, ok := p1Map[l.p1]; !ok && l.p1 != l.p2 {
			p1Map[l.p1] = true
		}
	}
	newLeaves := []node{}
	for _, l := range links {
		if _, ok := p1Map[l.p2]; !ok {
			// Then it's a leaf
			newLeaves = append(
				newLeaves,
				node{
					key:             l.p2,
					parentKey:       l.p1,
					parent:          nil,
					distToParent:    l.dist,
					children:        []*node{},
					descendantCount: 0,
				},
			)
		} else {
			remainingLinks = append(remainingLinks, l)
		}
	}
	for i, nl := range newLeaves {
		for j, ol := range leaves {
			if ol.parentKey == nl.key {
				newLeaves[i].children = append(newLeaves[i].children, &leaves[j])
				newLeaves[i].descendantCount = newLeaves[i].descendantCount + ol.descendantCount + 1
				leaves[j].parent = &newLeaves[i]
			}
		}
	}
	for _, ol := range leaves {
		if ol.parent == nil {
			newLeaves = append(newLeaves, ol)
		}
	}
	return buildClusterHierarchy(remainingLinks, newLeaves)
}

func condenseClusterTree(topNode *node, condensedTopCluster *cluster, minClusterSize int) cluster {
	if condensedTopCluster == nil {
		condensedTopCluster = &cluster{
			parent:           nil,
			points:           []int{},
			pointsDistParent: []float64{},
			children:         []*cluster{},
			lBirth:           0.0,
			lDeath:           0.0,
			stability:        0.0,
		}
	}
	for _, c := range topNode.children {
		if c.descendantCount >= minClusterSize {
			c2 := &cluster{
				parent:           condensedTopCluster,
				points:           []int{},
				pointsDistParent: []float64{},
				children:         []*cluster{},
				lBirth:           0.0,
				lDeath:           0.0,
				stability:        0.0,
			}
			c2.points = append(c2.points, c.key)
			c2.pointsDistParent = append(c2.pointsDistParent, c.distToParent)
			condensedTopCluster.children = append(condensedTopCluster.children, c2)
			condenseClusterTree(c, c2, minClusterSize)
		} else {
			condensedTopCluster.points = append(condensedTopCluster.points, c.key)
			condensedTopCluster.pointsDistParent = append(condensedTopCluster.pointsDistParent, c.distToParent)
		}
	}
	return *condensedTopCluster
}

func selectClusters(condensedTree *cluster, points [][]float64) []*node {
	return nil
}

func findClusters(dataPath string, savePath string, minClusterSize int, metric string, alpha float64, algorithm string, leafSize int, genMinSpanTree bool, clusterSelectionMethod string) {
	fmt.Println("Loading data")
	data := loadData(dataPath)

	// Compute the adjency matrix using mutual reachability
	mutualReachabilityArr := computeMutualReachability(minClusterSize, data)

	fmt.Println(mutualReachabilityArr)

	// Compute the min spanning tree using  the mutual reachability table created
	links := computeMinSpanningTree(mutualReachabilityArr)

	// Build the tree hierarchy structure of the clusters
	tree := buildClusterHierarchy(links, []node{})

	// Build the condensed tree
	condensedTree := condenseClusterTree(&tree, nil, minClusterSize)

	clusters := selectClusters(&condensedTree, data)

	fmt.Println("-----------------------")
	printTreeStructure(&tree)
	fmt.Println("-----------------------")
	printClusterStructure(&condensedTree)
	fmt.Println("-----------------------")
	for _, n := range clusters {
		printTreeStructure(n)
		fmt.Println("----------------------")
	}

	fmt.Println("Saving data")
	saveClusteringResult(savePath, data)
}
