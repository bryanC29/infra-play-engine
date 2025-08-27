package types

type Metrics struct {
	QPS1x  QPSMetrics `json:"1xQPS"`
	QPS15x QPSMetrics `json:"1.5xQPS"`
	QPS2x  QPSMetrics `json:"2xQPS"`
}

type QPSMetrics struct {
	Success      int
	Failed       int
	Latency      float64
	Availability string
}