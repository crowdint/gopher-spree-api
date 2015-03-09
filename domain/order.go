package domain

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/crowdint/gopher-spree-api/utils"
)

type Order struct {
	Id                  int64     `json:"id"`
	AdditionalTaxTotal  float64   `json:"additional_tax_total,string"`
	AdjustmentTotal     float64   `json:"adjustment_total,string"`
	BillAddressId       int64     `json:"-"`
	Channel             string    `json:"channel"`
	Currency            string    `json:"currency"`
	Email               string    `json:"email"`
	GuestToken          string    `json:"token"`
	IncludedTaxTotal    float64   `json:"included_tax_total,string"`
	ItemTotal           float64   `json:"item_total,string"`
	Number              string    `json:"number"`
	PaymentState        string    `json:"payment_state"`
	PaymentTotal        float64   `json:"payment_total,string"`
	ShipAddressId       int64     `json:"-"`
	ShipmentState       string    `json:"shipment_state"`
	ShipTotal           float64   `json:"ship_total,string" sql:"-"`
	SpecialInstructions string    `json:"special_instructions"`
	State               string    `json:"state"`
	TaxTotal            float64   `json:"tax_total,string" sql:"-"`
	Total               float64   `json:"total,string"`
	UserId              *int64    `json:"user_id"`
	CompletedAt         time.Time `json:"completed_at"`
	CreatedAt           time.Time `json:"created_at" sql:"-"`
	UpdatedAt           time.Time `json:"updated_at" sql:"-"`

	CheckoutSteps             []string     `json:"checkout_steps" sql:"-"` //TODO: implement
	DisplayAdditionalTaxTotal string       `json:"display_additional_tax_total" sql:"-"`
	DisplayIncludedTaxTotal   string       `json:"display_included_tax_total" sql:"-"`
	DisplayItemTotal          string       `json:"display_item_total" sql:"-"`
	DisplayTaxTotal           string       `json:"display_tax_total" sql:"-"`
	DisplayTotal              string       `json:"display_total" sql:"-"`
	DisplayShipTotal          string       `json:"display_ship_total" sql:"-"`
	Permissions               *Permissions `json:"permissions,omitempty" sql:"-"`
	Quantity                  int64        `json:"total_quantity" sql:"-"`

	BillAddress *Address     `json:"bill_address" sql:"-"`
	LineItems   *[]LineItem  `json:"line_items,omitempty" sql:"-"`
	Payments    []*Payment   `json:"payments" sql:"-"`
	ShipAddress *Address     `json:"ship_address" sql:"-""`
	Shipments   []*Shipment  `json:"shipments" sql:"-"`
	Adjustments []Adjustment `json:"adjustments" sql:"-"`

	ApprovedAt            time.Time `json:"-"`
	ApproverId            int64     `json:"-"`
	CanceledAt            time.Time `json:"-"`
	CancelerId            int64     `json:"-"`
	ConfirmationDelivered bool      `json:"-"`
	ConsideredRisky       bool      `json:"-"`
	CreatedBy             int64     `json:"-" sql:"-"`
	ItemCount             int64     `json:"-"`
	LastIpAddress         string    `json:"-"`
	PromoTotal            float64   `json:"-"`
	ShipmentTotal         float64   `json:"-"`
	ShippingMethodId      int64     `json:"-"`
	StateLockVersion      int64     `json:"-"`
	StoreId               int64     `json:"-"`
}

func (this Order) SpreeClass() string {
	return "Spree::Order"
}

func (this Order) TableName() string {
	return "spree_orders"
}

func (this Order) AdjustableCurrency() string {
	return this.Currency
}

func (this Order) AdjustableId() int64 {
	return this.Id
}

func (this *Order) AfterFind() (err error) {
	this.SetComputedValues()

	return
}

func (this *Order) Key() string {
	return fmt.Sprintf("%s/%d/%d", this.SpreeClass(), this.Id, this.UpdatedAt.Unix())
}

func (this *Order) KeyWithPrefix(prefix string) string {
	return fmt.Sprintf("%s/%s/%d/%d", this.SpreeClass(), prefix, this.Id, this.UpdatedAt.Unix())
}

func (this *Order) Marshal() ([]byte, error) {
	return json.Marshal(this)
}

func (this *Order) SetComputedValues() {
	this.TaxTotal = this.IncludedTaxTotal + this.AdditionalTaxTotal

	this.DisplayAdditionalTaxTotal = Monetize(this.AdditionalTaxTotal, this.Currency)
	this.DisplayIncludedTaxTotal = Monetize(this.IncludedTaxTotal, this.Currency)
	this.DisplayItemTotal = Monetize(this.ItemTotal, this.Currency)
	this.DisplayTaxTotal = Monetize(this.TaxTotal, this.Currency)
	this.DisplayTotal = Monetize(this.Total, this.Currency)
	this.DisplayShipTotal = Monetize(this.ShipTotal, this.Currency)
}

func (this *Order) Unmarshal(data []byte) error {
	return json.Unmarshal(data, this)
}
