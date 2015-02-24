package json

import (
	"time"
)

type Adjustment struct {
	Id             int64     `json:"id"`
	AdjustableId   int64     `json:"adjustable_id"`
	AdjustableType string    `json:"adjustable_type"`
	Amount         string    `json:"amount"`
	Eligible       bool      `json:"eligible"`
	Mandatory      *bool     `json:"mandatory"`
	Label          string    `json:"label"`
	SourceId       int64     `json:"source_id"`
	SourceType     string    `json:"source_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Computed
	DisplayAmount string `json:"display_amount"`
}

func (this Adjustment) TableName() string {
	return "spree_adjustments"
}
