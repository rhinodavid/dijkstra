#Dijkstra's Algorith

## Overview

This application computes the shortest paths between nodes in a directed graph using
a very simple implementation of [Dijkstra's Algorithm](https://en.wikipedia.org/wiki/Dijkstra's_algorithm).

## Input format

See `sample_graph.txt` for a sample input. The file is an adjacency list of a directed graph.
Each row consists of the label of a node followed by tuples representing edges that have
their tail in that node. The first element of the tuple is the ID of the head node of the edge
and the second element is the weight of that edge.

The graph in `sample_graph.txt` is represented below. The shortest path distances from
node 1 are shown in red.

<img src="diagram.png" width=350>

## Usage

Run the application with the first argument being the name of the graph adjacency list file,
the second argument the ID of the starting node, and subsequent arguments the IDs of nodes
you would like to view the shortest path distance of.

For example, run `go run main.go sample_graph.txt 1 3 5` and you should see:

```
The shortest distance for Node 3 is: 16
The shortest distance for Node 5 is: 17
```

If a shortest distance is displayed as -1, that means that there was no path between
the starting node and the selected node.
