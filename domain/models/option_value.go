package models

import "time"

type OptionValue struct {
	Id           int64
	Position     int64
	Name         string
	Presentation string
	OptionTypeId int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (this OptionValue) TableName() string {
	return "spree_option_values"
}
