package types

type Graph struct {
	Nodes       []Node
	Index       map[string]int // nodeID -> idx
	Out         [][]int        // adjacency forward
	In          [][]int        // adjacency reverse
	Entry, Exit int            // indices of special nodes
	Topo        []int          // topological order (filled after DAG check)
}

type Path struct {
	Nodes []int // indices including entry and exit
}