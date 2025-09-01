package types

type Problem struct {
	SubmissionID string `json:"submissionId" validate:"required"` // Magic number to track a unique submission
	BaseQPS      int    `json:"baseQps" validate:"required"`      // Baseline QPS to simulate
	Design       Design `json:"design" validate:"required"`
}