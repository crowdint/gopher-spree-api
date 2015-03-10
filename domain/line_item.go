package domain

import (
	"time"

	. "github.com/crowdint/gopher-spree-api/utils"
)

type LineItem struct {
	Id                 int64     `json:"id"`
	AdditionalTaxTotal float64   `json:"-"`
	AdjustmentTotal    float64   `json:"-"`
	CostPrice          float64   `json:"-"`
	Currency           string    `json:"-"`
	IncludedTaxTotal   float64   `json:"-"`
	PreTaxAmount       float64   `json:"-"`
	OrderId            int64     `json:"-"`
	Price              float64   `json:"price,string"`
	PromoTotal         float64   `json:"-"`
	Quantity           int64     `json:"quantity"`
	TaxCategoryId      int64     `json:"-"`
	VariantId          int64     `json:"-"`
	CreatedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`

	Amount              float64 `json:"-" sql:"-"`
	DisplayAmount       string  `json:"display_amount" sql:"-"`
	FinalAmount         float64 `json:"total,string" sql:"-"`
	SingleDisplayAmount string  `json:"single_display_amount" sql:"-"`

	Adjustments []Adjustment `json:"adjustments"`
	Variant     *Variant     `json:"variant" sql:"-"`
}

func (this LineItem) AdjustableId() int64 {
	return this.Id
}

func (this LineItem) AdjustableCurrency() string {
	return this.Currency
}

func (this *LineItem) AfterFind() (err error) {
	this.Amount = this.Price * float64(this.Quantity)
	this.FinalAmount = this.Amount + this.AdjustmentTotal //TODO: this should match spree api (rounded).

	this.DisplayAmount = Monetize(this.Amount, this.Currency)
	this.SingleDisplayAmount = Monetize(this.Price, this.Currency)
	return
}

func (this LineItem) TableName() string {
	return "spree_line_items"
}

func (this LineItem) SpreeClass() string {
	return "Spree::LineItem"
}
