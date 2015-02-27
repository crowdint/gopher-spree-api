package domain

import "time"

type OptionValue struct {
	Id                     int64     `json:"id"`
	Position               int64     `json:"-"`
	Name                   string    `json:"name"`
	Presentation           string    `json:"presentation"`
	OptionTypeId           int64     `json:"option_type_id"`
	OptionTypePresentation string    `json:"option_type_presentation" sql:"-"`
	OptionTypeName         string    `json:"option_type_name" sql:"-"`
	CreatedAt              time.Time `json:"-"`
	UpdatedAt              time.Time `json:"-"`
	VariantId              int64     `json:"-" sql:"-"`
}

func (this OptionValue) TableName() string {
	return "spree_option_values"
}
