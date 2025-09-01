package types

type Metrics struct {
	QPS1x  QPSMetrics `json:"1xQPS"`
	QPS15x QPSMetrics `json:"1.5xQPS"`
	QPS2x  QPSMetrics `json:"2xQPS"`

	AvgSuccess      float64 `json:"average success"`
	AvgFailed       float64 `json:"average failed"`
	AvgLatency      float64 `json:"average latency"`
	AvgAvailability float64 `json:"average availability"`
}

type QPSMetrics struct {
	Success      int
	Failed       int
	Latency      float64
	Availability float64
}