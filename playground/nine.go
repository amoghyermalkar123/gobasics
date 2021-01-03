package playground

// implementation of Dijkstra' algorithm

type Dnode struct {
	value int
}

type Dedge struct {
	source      *Dnode
	destination *Dnode
	weight      int
}

var set []Dnode

func pickLeastWeightedEdge(source Dnode, graph []Dedge) {
	// d := &Dnode{}
	// allNodes := d.initNodes()
	costMatrix := make([]int, 0)
	for item := range graph {
		// find all neighbors of the source node
		if graph[item].source.value == source.value {
			costMatrix = append(costMatrix, costMatrix[source.value])
		}
	}
}

func (d *Dnode) initNodes() (costMatrix [9]int) {
	for i := 0; i < 9; i++ {
		d.value = i
	}
	costMatrix = [9]int{}
	return costMatrix
}
func updateNeighbors() {

}

// func BuildDgraph() {
// 	graph := GetGraph()
// 	pickLeastWeightedEdge(graph)
// }
