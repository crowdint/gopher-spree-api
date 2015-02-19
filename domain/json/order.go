package json

import (
	"time"
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
	ShipTotal           float64   `json:"ship_total,string"`
	SpecialInstructions string    `json:"special_instructions"`
	State               string    `json:"state"`
	TaxTotal            float64   `json:"tax_total,string"`
	Total               float64   `json:"total,string"`
	UserId              *int64    `json:"user_id"`
	CompletedAt         time.Time `json:"completed_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`

	//Computed
	CheckoutSteps             []string     `json:"checkout_steps"`               //TODO: implement
	DisplayAdditionalTaxTotal string       `json:"display_additional_tax_total"` //TODO: implement
	DisplayIncludedTaxTotal   string       `json:"display_included_tax_total"`   //TODO: implement
	DisplayItemTotal          string       `json:"display_item_total"`           //TODO: implement
	DisplayTaxTotal           string       `json:"display_tax_total"`            //TODO: implement
	DisplayTotal              string       `json:"display_total"`                //TODO: implement
	DisplayShipTotal          string       `json:"display_ship_total"`           //TODO: implement
	Permissions               *Permissions `json:"permissions,omitempty"`
	Quantity                  int64        `json:"total_quantity"`

	// Associations
	BillAddress *Address    `json:"bill_address,omitempty"`
	LineItems   *[]LineItem `json:"line_items,omitempty"`
	Payments    []Payment   `json:"payments,omitempty"`
	ShipAddress *Address    `json:"ship_address,omitempty"`
	Shipments   []Shipment  `json:"shipments,omitempty"`
}
