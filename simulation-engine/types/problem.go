package types

type Problem struct {
	UserID               string  `json:"userId" validate:"required"`          // ID of the user submitting the design
	SubmissionID         string  `json:"submissionId" validate:"required"`    // Magic number to track a unique submission
	BaseQPS              int     `json:"baseQps" validate:"required"`         // Baseline QPS to simulate
	RequiredAvailability float64 `json:"reqAvailability" validate:"required"` // Target availability (e.g., 0.99)
	RequiredLatencyMS    int     `json:"reqLatencyMs" validate:"required"`    // Max allowed average latency in milliseconds
	Design               Design  `json:"design" validate:"required"`
}