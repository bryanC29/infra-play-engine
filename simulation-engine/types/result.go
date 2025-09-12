package types

type Result struct {
	SubmissionID string  `json:"submissionID"`
	BaseQPS      int     `json:"baseQps"`
	Nodes        int     `json:"totalNodes"`
	Metrics      Metrics `json:"metrics"`
}