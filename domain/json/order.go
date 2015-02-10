package json

import (
	"fmt"
	"strconv"
	"time"
)

type Order struct {
	Id                        int64        `json:"id"`
	Number                    string       `json:"number"`
	ItemTotal                 float64      `json:"item_total,string"`
	Total                     float64      `json:"total,string"`
	ShipTotal                 float64      `json:"ship_total,string"`
	State                     string       `json:"state"`
	AdjustmentTotal           float64      `json:"adjustment_total,string"`
	UserId                    *int64       `json:"user_id"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt                 time.Time    `json:"updated_at"`
	CompletedAt               time.Time    `json:"completed_at"`
	PaymentTotal              float64      `json:"payment_total,string"`
	ShipmentState             string       `json:"shipment_state"`
	PaymentState              string       `json:"payment_state"`
	Email                     string       `json:"email"`
	SpecialInstructions       string       `json:"special_instructions"`
	Channel                   string       `json:"channel"`
	IncludedTaxTotal          float64      `json:"included_tax_total,string"`
	AdditionalTaxTotal        float64      `json:"additional_tax_total,string"`
	DisplayIncludedTaxTotal   string       `json:"display_included_tax_total"`
	DisplayAdditionalTaxTotal string       `json:"display_additional_tax_total"`
	TaxTotal                  float64      `json:"tax_total,string"`
	Currency                  string       `json:"currency"`
	DisplayItemTotal          string       `json:"display_item_total"`
	Quantity                  int64        `json:"total_quantity"`
	DisplayTotal              string       `json:"display_total"`
	DisplayShipTotal          string       `json:"display_ship_total"`
	DisplayTaxTotal           string       `json:"display_tax_total"`
	GuestToken                string       `json:"token"`
	CheckoutSteps             []string     `json:"checkout_steps"`
	Permissions               *Permissions `json:"permissions,omitempty"`
	BillAddress               *Address     `json:"bill_address,omitempty"`
	BillAddressId             int64        `json:"-"`
	ShipAddress               *Address     `json:"ship_address,omitempty"`
	ShipAddressId             int64        `json:"-"`
	LineItems                 *[]LineItem  `json:"line_items,omitempty"`
	Payments                  []Payment    `json:"payments,omitempty"`
	Shipments                 []Shipment   `json:"shipments,omitempty"`
}

type Address struct {
	Id               int64    `json:"id"`
	Firstname        string   `json:"firstname"`
	Lastname         string   `json:"lastname"`
	FullName         string   `json:"full_name"`
	Address1         string   `json:"address1"`
	Address2         string   `json:"address2"`
	City             string   `json:"city"`
	Zipcode          string   `json:"zipcode"`
	Phone            string   `json:"phone"`
	Company          string   `json:"company"`
	AlternativePhone string   `json:"alternative_phone"`
	CountryId        int64    `json:"country_id"`
	StateId          int64    `json:"state_id"`
	StateName        string   `json:"state_name"`
	StateText        string   `json:"state_text"`
	Country          *Country `json:"country"`
	State            *State   `json:"state"`
}

func (this *Address) AfterFind() (err error) {
	this.FullName = fmt.Sprintf("%s %s", this.Firstname, this.Lastname)
	return
}

func (this Address) TableName() string {
	return "spree_addresses"
}

type Adjustment struct {
	Id             int64     `json:"id"`
	SourceType     string    `json:"source_type"`
	SourceId       int64     `json:"source_id"`
	AdjustableType string    `json:"adjustable_type"`
	AdjustableId   int64     `json:"adjustable_id"`
	Amount         string    `json:"amount"`
	Label          string    `json:"label"`
	Mandatory      bool      `json:"mandatory"`
	Elegible       bool      `json:"elegible"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DisplayAmount  string    `json:"display_amount"`
}

type Country struct {
	Id      int64  `json:"id"`
	IsoName string `json:"iso_name"`
	Iso     string `json:"iso"`
	Iso3    string `json:"iso3"`
	Name    string `json:"name"`
	Numcode int64  `json:"numcode"`
}

func (this Country) TableName() string {
	return "spree_countries"
}

type LineItem struct {
	Id                  int64        `json:"id"`
	Quantity            int64        `json:"quantity"`
	Price               float64      `json:"price,string"`
	VariantId           int64        `json:"variant_id"`
	SingleDisplayAmount string       `json:"single_display_amount"`
	DisplayAmount       string       `json:"display_amount"`
	Total               string       `json:"total"`
	Variant             *Variant     `json:"variant"`
	Adjustments         []Adjustment `json:"adjustments"`

	Amount int64 `json:"-"`

	OrderId         int64   `json:"-"`
	AdjustmentTotal float64 `json:"-"`
}

func (this *LineItem) AfterFind() (err error) {
	// Compute Total
	amount := this.Price * float64(this.Quantity)
	finalAmount := amount + this.AdjustmentTotal
	this.Total = strconv.FormatFloat(finalAmount, 'f', 2, 64)

	return
}

func (this LineItem) TableName() string {
	return "spree_line_items"
}

type State struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	CountryId int64  `json:"country_id"`
}

func (this State) TableName() string {
	return "spree_states"
}

type Payment struct {
	Id              int64         `json:"id"`
	SourceType      string        `json:"source_type"`
	SourceId        int64         `json:"source_id"`
	Amount          string        `json:"amount"`
	DisplayAmount   string        `json:"display_amount"`
	PaymentMethodId int64         `json:"payment_method_id"`
	ResponseCode    string        `json:"response_code"`
	State           string        `json:"state"`
	AVSResponse     string        `json:"avs_response"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	Source          PaymentSource `json:"source"`
}

type PaymentMethod struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type PaymentSource struct {
	Id                       int64  `json:"id"`
	Month                    string `json:"month"`
	Year                     string `json:"year"`
	CCType                   string `json:"cc_type"`
	LastDigits               string `json:"last_digits"`
	Name                     string `json:"name"`
	GatewayCustomerProfileId string `json:"gateway_customer_profile_id"`
	GatewayPaymentProfileId  string `json:"gateway_payment_profile_id"`
}

type Permissions struct {
	CanUpdate *bool `json:"can_update,omitempty"` // user.HasRole("admin") || (order.UserId == user.Id)
}

type Shipment struct {
	Id                   int64            `json:"id"`
	Tracking             string           `json:"tracking"`
	Number               string           `json:"number"`
	Cost                 string           `json:"cost"`
	ShippedAt            *time.Time       `json:"shipped_at"`
	State                string           `json:"state"`
	OrderId              string           `json:"order_id"`
	StockLocationName    string           `json:"stock_location_name"`
	ShippingRates        []ShippingRate   `json:"shipping_rates"`
	SelectedShippingRate ShippingRate     `json:"selected_shipping_rate"`
	ShippingMethods      []ShippingMethod `json:"shipping_methods"`
	Manifest             ShipmentManifest `json:"shipment_manifest"`
	Adjustments          []Adjustment     `json:"adjustments"`
}

type ShipmentManifest struct {
	Quantity  int64            `json:"quantity"`
	States    map[string]int64 `json:"states"`
	VariantId int64            `json:"variant_id"`
}

type ShippingCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ShippingMethod struct {
	Id                 int64              `json:"id"`
	Code               string             `json:"code"`
	Name               string             `json:"name"`
	Zones              []ShippingZone     `json:"zones"`
	ShippingCategories []ShippingCategory `json:"shipping_categories"`
}

type ShippingRate struct {
	Id                 int64   `json:"id"`
	Name               string  `json:"name"`
	Cost               string  `json:"cost"`
	Selected           bool    `json:"selected"`
	ShippingMethodId   float64 `json:"shipping_method_id"`
	ShippingMethodCode string  `json:"shipping_method_code"`
	DisplayCost        string  `json:"display_cost"`
}

type ShippingZone struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
