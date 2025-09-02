package simulator

import (
	"math"
	"simengine/types"
)

func SimulateGlobal(g *types.Graph, entryLoad float64) *types.QPSMetrics {
	const QPSPerCPU = 250.0
	const BaseLatencyPerNodeMs = 1.5

	n := len(g.Nodes)
	load := make([]float64, n)
	load[g.Entry + 1] = entryLoad

	var totalFail, totalLatency float64

	for _, u := range g.Topo {

		if u == g.Entry || u == g.Exit {
			continue
		}

		incoming := load[u]
		if incoming == 0 {
			continue
		}

		node := g.Nodes[u]
		capacity := float64(node.Resources.CPU) * QPSPerCPU * float64(node.Resources.Replicas)

		processed := incoming
		if processed > capacity {
			over := processed - capacity
			processed = capacity
			totalFail += over
		}

		totalLatency += processed * BaseLatencyPerNodeMs

		if len(g.Out[u]) > 0 {
			share := processed / float64(len(g.Out[u]))
			for _, v := range g.Out[u] {
				load[v] += share
			}
		}
	}

	latency := totalLatency / entryLoad
	availability := (load[g.Exit] / entryLoad) * 100
	
	return &types.QPSMetrics {
    	Availability: math.Round(availability * 100) / 100,
    	Latency:      math.Round(latency * 100) / 100,
    	Failed:       int(totalFail),
    	Success:      int(load[g.Exit]),
	}
}