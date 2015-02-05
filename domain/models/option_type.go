package models

import "time"

type OptionType struct {
	Id           int64
	Name         string
	Presentation string
	Position     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductId    int64
}

func (this OptionType) TableName() string {
	return "spree_option_types"
}
