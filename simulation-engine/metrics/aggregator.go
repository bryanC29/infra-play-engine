package metrics

import (
	"math"
	"simengine/types"
)

func MetricsAggregator(m *types.Metrics) (*types.Metrics, error) {
	totalSucc := m.QPS1x.Success + m.QPS15x.Success + m.QPS2x.Success
	avgSucc := totalSucc / 3
	m.AvgSuccess = float64(avgSucc)	
	
	totalFail := m.QPS1x.Failed + m.QPS15x.Failed + m.QPS2x.Failed
	avgFail := totalFail / 3
	m.AvgFailed = float64(avgFail)	
	
	totalLatency := m.QPS1x.Success + m.QPS15x.Success + m.QPS2x.Success
	avgLatency := totalLatency / 3
	m.AvgLatency = float64(avgLatency)	
	
	totalAvail := m.QPS1x.Availability + m.QPS15x.Availability + m.QPS2x.Availability
	avgAvail := totalAvail / 3
	m.AvgAvailability = math.Round(avgAvail * 100) / 100
	return m, nil
}