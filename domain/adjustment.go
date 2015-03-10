package domain

import (
	"time"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	. "github.com/crowdint/gopher-spree-api/utils"
)

type Adjustment struct {
	Id             int64     `json:"id"`
	AdjustableId   int64     `json:"adjustable_id"`
	AdjustableType string    `json:"adjustable_type"`
	Amount         float64   `json:"amount,string"`
	Eligible       bool      `json:"eligible"`
	Mandatory      *bool     `json:"mandatory"`
	Label          string    `json:"label"`
	SourceId       int64     `json:"source_id"`
	SourceType     string    `json:"source_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	DisplayAmount string `json:"display_amount" sql:"-"`

	Adjustable Adjustable `json:"-" sql:"-"`
}

func (this *Adjustment) AfterFind() (err error) {
	this.DisplayAmount = Monetize(this.Amount, this.Currency())

	return
}

func (this *Adjustment) Currency() string {
	if this.Adjustable != nil {
		return this.Adjustable.AdjustableCurrency()
	}

	return spree.Get(spree.CURRENCY)
}

func (this Adjustment) TableName() string {
	return "spree_adjustments"
}
