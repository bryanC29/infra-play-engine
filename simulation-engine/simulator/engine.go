package simulator

import (
	"fmt"
	"simengine/types"
)

func EngineRun(design types.Design) (types.Metrics, error) {
	qps := float64(100)
	var res types.Metrics
	result, err := BuildGraph(design)
	disConn := ContainsIsolatedNodes(design)

	if disConn {
		return types.Metrics{}, fmt.Errorf("graph is disconnected")
	}

	if err != nil {	
		return types.Metrics{}, fmt.Errorf("an error occured %w", err)
	}

	res.QPS1x = SimulateGlobal(result, qps)
	res.QPS15x = SimulateGlobal(result, qps * 1.5)
	res.QPS2x = SimulateGlobal(result, qps * 2)
	
	return res, nil
}