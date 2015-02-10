package models

import "time"

type OptionValue struct {
	Id                     int64
	Position               int64
	Name                   string
	Presentation           string
	OptionTypeId           int64
	OptionTypePresentation string
	OptionTypeName         string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	VariantId              int64
}

func (this OptionValue) TableName() string {
	return "spree_option_values"
}
