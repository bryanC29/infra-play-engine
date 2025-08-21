package simulator

import (
	"log"
	"simengine/types"
)

func EngineRun(design types.Design) {
	qps := float64(100)
	result, _ := BuildGraph(design)
	
	log.Print("Node index: ", result.Index)
	log.Print("Topo sort: ", result.Topo)
	log.Print("Out: ", result.Out)
	log.Print("Disconnected: ", ContainsIsolatedNodes(design))

	log.Print("1x QPS")
	SimulateGlobal(result, qps)
	
	log.Print("1.5x QPS")
	SimulateGlobal(result, qps * 1.5)
	
	log.Print("2x QPS")
	SimulateGlobal(result, qps * 2)
}