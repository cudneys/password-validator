package models

import "github.com/nbutton23/zxcvbn-go/scoring"

type Response struct {
	Error      string                  `json:"error,omitempty"`
	IsValid    bool                    `json:"is_valid""`
	Assessment scoring.MinEntropyMatch `json:"summary,omitempty"`
}
