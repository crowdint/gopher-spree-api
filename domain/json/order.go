package json

import (
	"time"
)

type Order struct {
	Id                        int64       `json:"id"`
	Number                    string      `json:"number"`
	ItemTotal                 string      `json:"item_total"`
	Total                     string      `json:"total"`
	ShipTotal                 string      `json:"ship_total"`
	State                     string      `json:"state"`
	AdjustmentTotal           string      `json:"adjustment_total"`
	UserId                    *int64      `json:"user_id"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
	CompletedAt               time.Time   `json:"completed_at"`
	PaymentTotal              string      `json:"payment_total"`
	ShipmentState             string      `json:"shipment_state"`
	PaymentState              string      `json:"payment_state"`
	Email                     string      `json:"email"`
	SpecialInstructions       string      `json:"special_instructions"`
	Channel                   string      `json:"channel"`
	IncludedTaxTotal          string      `json:"included_tax_total"`
	AdditionalTaxTotal        string      `json:"additional_tax_total"`
	DisplayIncludedTaxTotal   string      `json:"display_included_tax_total"`
	DisplayAdditionalTaxTotal string      `json:"display_additional_tax_total"`
	TaxTotal                  string      `json:"tax_total"`
	Currency                  string      `json:"currency"`
	DisplayItemTotal          string      `json:"display_item_total"`
	TotalQuantity             int64       `json:"total_quantity"`
	DisplayTotal              string      `json:"display_total"`
	DisplayShipTotal          string      `json:"display_ship_total"`
	DisplayTaxTotal           string      `json:"display_tax_total"`
	Token                     string      `json:"token"`
	CheckoutSteps             []string    `json:"checkout_steps"`
	Permissions               Permissions `json:"permissions"`
	BillAddress               Address     `json:"bill_address"`
	ShipAddress               Address     `json:"ship_address"`
	LineItems                 []LineItem  `json:"line_items"`
}

type Address struct {
	Id               int64   `json:"id"`
	FirstName        string  `json:"firstname"`
	LastName         string  `json:"lastname"`
	FullName         string  `json:"full_name"`
	Address1         string  `json:"address1"`
	Address2         string  `json:"address2"`
	City             string  `json:"city"`
	ZipCode          string  `json:"zipcode"`
	Phone            string  `json:"phone"`
	Company          string  `json:"company"`
	AlternativePhone string  `json:"alternative_phone"`
	CountryId        int64   `json:"country_id"`
	StateId          int64   `json:"state_id"`
	StateName        string  `json:"state_name"`
	StateText        string  `json:"state_text"`
	Country          Country `json:"country"`
	State            State   `json:"state"`
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
	ISOName string `json:"iso_name"`
	ISO     string `json:"iso"`
	ISO3    string `json:"iso3"`
	Name    string `json:"name"`
	NumCode int64  `json:"numcode"`
}

type LineItem struct {
	Id                  int64        `json:"id"`
	Quantity            int64        `json:"quantity'`
	Price               string       `json:"price"`
	VariantId           int64        `json:"variant_id"`
	SingleDisplayAmount string       `json:"singe_display_amount"`
	DisplayAmount       string       `json:"display_amount"`
	Total               string       `json:"total"`
	Variant             Variant      `json:"variant"`
	Adjustments         []Adjustment `json:"adjustments"`
}

type State struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	CountryId int64  `json:"country_id"`
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
	Id                       int64   `json:"id"`
	Month                    string  `json:"month"`
	Year                     string  `json:"year"`
	CCType                   string  `json:"cc_type"`
	LastDigits               string  `json:"last_digits"`
	Name                     string  `json:"name"`
	GatewayCustomerProfileId string  `json:"gateway_customer_profile_id"`
	GatewayPaymentProfileId  *string `json:"gateway_payment_profile_id"`
}

type Permissions struct {
	CanUpdate bool `json:"can_update"`
}
