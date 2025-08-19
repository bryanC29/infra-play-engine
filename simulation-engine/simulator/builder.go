package simulator

import (
	"errors"
	"fmt"
	"log"
	"simengine/types"
	"sort"
)

func BuildGraph(d types.Design) (*types.Graph, error) {
	if len(d.Nodes) == 0 {
		return nil, errors.New("no nodes provided")
	}
	
	g := &types.Graph{Index: make(map[string]int)}
	g.Nodes = make([]types.Node, len(d.Nodes))

	for i, n := range d.Nodes {
		if _, exist := g.Index[n.ID]; exist {
			return nil, fmt.Errorf("duplicate node id: %s", n.ID)
		}
		g.Index[n.ID] = i
		g.Nodes[i] = n
	}

	log.Printf("hello")
	log.Printf("%v", g.Index)

	entryIdx, entryExist := g.Index["entry"]
	if !entryExist {
		return nil, errors.New("no entry node found")
	}
	
	exitIdx, exitExist := g.Index["exit"]
	if !exitExist {
		return nil,errors.New("no exit node found")
	}
	g.Entry, g.Exit = entryIdx, exitIdx

	n := len(g.Nodes)
	g.Out = make([][]int, n)
	g.In = make([][]int, n)

	seenEdge := make(map[[2]int]struct{})
	
	for _, c := range d.Connections {
		from, ok1 := g.Index[c.From]
		to, ok2 := g.Index[c.To]
		
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("edge references unknown node: %s -> %s", c.From, c.To)
		}
		
		if from == to {
			return nil, fmt.Errorf("self-loop not allowed: %s", c.From)
		}
		
		key := [2]int{from, to}
		
		if _, dup := seenEdge[key]; dup {
			return nil, fmt.Errorf("duplicate edge: %s -> %s", c.From, c.To)
		}
		
		seenEdge[key] = struct{}{}
		g.Out[from] = append(g.Out[from], to)
		g.In[to] = append(g.In[to], from)
	}

	for i := range g.Out {
		sort.Ints(g.Out[i])
	}
	for i := range g.In {
		sort.Ints(g.In[i])
	}

	order, ok := topoKahn(g.Out, g.In)
	if !ok {
		return nil, errors.New("graph contains cycle")
	}

	g.Topo = order

	if !reachable(g.Out, g.Entry, g.Exit) {
		return nil, errors.New("exit is unreachable from entry")
	}

	return g, nil
}

func topoKahn(out [][]int, in [][]int) ([]int, bool) {
	n := len(out)
	inDeg := make([]int, n)
	q := make([]int, 0, n)
	order := make([]int, 0, n)

	for v := range n {
		inDeg[v] = len(in[v])
	}

	for v := range n {
		if inDeg[v] == 0 {
			q = append(q, v)
		}
	}

	for i := range len(q) {
		v := q[i]
		order = append(order, v)
		for _, w := range out[v] {
			inDeg[w]--
			if inDeg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	if len(order) != n {
		return nil, false
	}
	
	return order, true
}

func reachable(out [][]int, src, dst int) bool {
	n := len(out)
	vis := make([]bool, n)
	q := []int{src}
	vis[src] = true
	
	for i := range len(q) {
		v := q[i]
		
		if v == dst {
			return true
		}
		
		for _, w := range out[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, w)
			}
		}
	}
	
	return false
}