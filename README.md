**Project is not in working condition**

# HDBSCAN

A go implementation of HDBSCAN based on [the original paper](https://link.springer.com/chapter/10.1007/978-3-642-37456-2_14) and [this excellent writeup](https://hdbscan.readthedocs.io/en/latest/how_hdbscan_works.html) by the author of the scikit-learn module for HDBSCAN.

## How does it work

The executable takes the following parameters:

```
Usage of ./hdbscan:
  -algorithm string
    	Which algorithm to use. Choices are best, generic, prims_kdtree, prims_balltree, boruvka_kdtree, boruvka_balltree (default "best")
  -alpha float
    	Alpha value (default 1)
  -cluster_selection_method string
    	Method to select clusters from the condensed tree (default "eom")
  -data_path string
    	Data CSV file path (default "data.csv")
  -gen_min_span_tree
    	Wheter to generate the minimum spanning tree with regard to mutual reachability
  -leaf_size int
    	Number of leaf node in the tree (when using a tree-based method) (default 40)
  -metric string
    	Distance metric used for the clustering (default "euclidean")
  -min_cluster_size int
    	Minimal cluster size (default 5)
  -save_path string
    	Data CSV save path (default "save.csv")
```

As you might have figured out, these are directly taken from the Python implementation. For clarification on their respective usage, please see [the module's documentation](https://hdbscan.readthedocs.io/en/latest/api.html#hdbscan).

It is worth noting that most of these parameters are useless at the moment and are subject to change.

## Implementation details

This is a non-optimal implementation in its current form. The algorithm used to build the MST is O(n^2) and I am fairly sure some computations are wasted here and there.

Notably, the implementation uses float64 instead of the smaller float32. This is mostly due to golang's "default" being float64 and not because the precision is needed.

## Long term goals

Ideally I would like to turn this into an actual golang module instead of a CLI executable in the future.
