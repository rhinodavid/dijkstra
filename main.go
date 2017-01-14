package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rhinodavid/dijkstra/graph"
)

func main() {
	g := graph.New()

	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the graph adjacency list")
	}
	f, err := os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		iS := strings.Fields(line)
		if len(iS) < 1 {
			panic(fmt.Sprintf("Bad line in graph file: %v", iS))
		}
		nID, err1 := strconv.Atoi(iS[0])
		if err1 != nil {
			panic(err1)
		}
		n := g.AddNode(nID)
		for i := 1; i < len(iS); i++ {
			eS := strings.Split(iS[i], ",")
			v, _ := strconv.Atoi(eS[0])
			len, _ := strconv.Atoi(eS[1])
			n.AddEdge(v, len)
		}
	}
	sS := flag.Args()[1]
	s, sErr := strconv.Atoi(sS)
	if sErr != nil {
		log.Fatal("The second argument must be the id of the node to compute shortest paths from")
	}
	g.ComputeShortestDistances(s)
	for _, id := range os.Args[3:] {
		idInt, _ := strconv.Atoi(id)
		g.ShowShortestDistance(idInt)
	}
}
