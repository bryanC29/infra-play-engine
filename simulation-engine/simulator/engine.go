package simulator

import (
	"fmt"
	"simengine/metrics"
	"simengine/types"
)

func EngineRun(problem types.Problem) (*types.Result, error) {
	qps := float64(problem.BaseQPS)
	var metric = &types.Metrics{}
	result, err := BuildGraph(problem.Design)
	disConn := ContainsIsolatedNodes(problem.Design)

	if disConn {
		return &types.Result{}, fmt.Errorf("graph is disconnected")
	}

	if err != nil {	
		return &types.Result{}, fmt.Errorf("an error occured %w", err)
	}
	
	metric.QPS1x = *SimulateGlobal(result, qps)
	metric.QPS15x = *SimulateGlobal(result, qps * 1.5)
	metric.QPS2x = *SimulateGlobal(result, qps * 2)

	res, err := metrics.MetricsAggregator(metric)
	
	if err != nil {	
		return &types.Result{}, fmt.Errorf("an error occured %w", err)
	}
	
	ans, errRes := metrics.ResultAggregator(*result, problem, *res)
	
	if errRes != nil {	
		return &types.Result{}, fmt.Errorf("an error occured %w", err)
	}

	return ans, nil
}