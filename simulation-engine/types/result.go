package types

type Result struct {
	SubmissionID string  `json:"SubmissionID"`
	BaseQPS      int     `json:"base QPS"`
	Nodes        int     `json:"total nodes"`
	Metrics      Metrics `json:"metrics"`
}