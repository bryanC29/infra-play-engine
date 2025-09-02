package metrics

import (
	"simengine/types"
)

func ResultAggregator(g types.Graph, p types.Problem, m types.Metrics) (*types.Result, error) {
	res := &types.Result{}

	res.Metrics = m
	res.SubmissionID = p.SubmissionID
	res.BaseQPS = p.BaseQPS
	res.Nodes = len(g.Nodes)

	return res, nil
}