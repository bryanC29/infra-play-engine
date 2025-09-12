package types

type Metrics struct {
	QPS1x  QPSMetrics `json:"qps1x"`
	QPS15x QPSMetrics `json:"qps15x"`
	QPS2x  QPSMetrics `json:"qps2x"`

	AvgSuccess      float64 `json:"avgSuccess"`
	AvgFailed       float64 `json:"avgFailed"`
	AvgLatency      float64 `json:"avgLatency"`
	AvgAvailability float64 `json:"avgAvail"`
}

type QPSMetrics struct {
	Success      int     `json:"success"`
	Failed       int     `json:"failed"`
	Latency      float64 `json:"latency"`
	Availability float64 `json:"avail"`
}