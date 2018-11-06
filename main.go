package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
		HDBSCAN(
			min_cluster_size=5,
			metric='euclidean',
			alpha=1.0,
			algorithm='best',
			leaf_size=40,
			gen_min_span_tree=False,
			cluster_selection_method='eom')
	*/
	dataFilePathPtr := flag.String("data_path", "data.csv", "Data CSV file path")
	clusteringDataSavePathPtr := flag.String("save_path", "save.csv", "Data CSV save path")
	minClusterSizePtr := flag.Int("min_cluster_size", 5, "Minimal cluster size")
	metricPtr := flag.String("metric", "euclidean", "Distance metric used for the clustering")
	alphaPtr := flag.Float64("alpha", 1.0, "Alpha value")
	algorithmPtr := flag.String("algorithm", "best", "Which algorithm to use. Choices are best, generic, prims_kdtree, prims_balltree, boruvka_kdtree, boruvka_balltree")
	leafSizePtr := flag.Int("leaf_size", 40, "Number of leaf node in the tree (when using a tree-based method)")
	genMinSpanTreePtr := flag.Bool("gen_min_span_tree", false, "Wheter to generate the minimum spanning tree with regard to mutual reachability")
	clusterSelectionMethodPtr := flag.String("cluster_selection_method", "eom", "Method to select clusters from the condensed tree")

	flag.Parse()

	fmt.Println("Data:", *dataFilePathPtr)
	fmt.Println("Save:", *clusteringDataSavePathPtr)
	fmt.Println("mCS:", *minClusterSizePtr)
	fmt.Println("m:", *metricPtr)
	fmt.Println("a:", *alphaPtr)
	fmt.Println("a:", *algorithmPtr)
	fmt.Println("lS:", *leafSizePtr)
	fmt.Println("gMST:", *genMinSpanTreePtr)
	fmt.Println("cSM", *clusterSelectionMethodPtr)

	findClusters(
		*dataFilePathPtr,
		*clusteringDataSavePathPtr,
		*minClusterSizePtr,
		*metricPtr,
		*alphaPtr,
		*algorithmPtr,
		*leafSizePtr,
		*genMinSpanTreePtr,
		*clusterSelectionMethodPtr,
	)
}
