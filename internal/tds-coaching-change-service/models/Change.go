package models

import "github.com/ocrosby/soccer/pkg/common"

// Change represents a coaching change within a program.
// swagger:model
type Change struct {
	Program           string        `json:"program_name"`
	ProgramId         int           `json:"program_id"`
	ProgramUrl        string        `json:"program_url"`
	DepartingCoach    string        `json:"departing_coach"`
	DepartingCoachUrl string        `json:"departing_coach_url"`
	NewCoach          string        `json:"new_coach"`
	NewCoachUrl       string        `json:"new_coach_url"`
	Gender            common.Gender `json:"program_gender"`
}
