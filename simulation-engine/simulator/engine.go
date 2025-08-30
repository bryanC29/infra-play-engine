package simulator

import (
	"fmt"
	"simengine/types"
)

func EngineRun(problem types.Problem) (types.Metrics, error) {
	qps := float64(problem.BaseQPS)
	var res types.Metrics
	result, err := BuildGraph(problem.Design)
	disConn := ContainsIsolatedNodes(problem.Design)

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