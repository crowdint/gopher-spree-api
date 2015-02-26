package domain

import "time"

type OptionType struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Presentation string `json:"presentation"`
	Position     int64  `json:"position"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ProductId int64     `json:"-"`
}

func (this OptionType) TableName() string {
	return "spree_option_types"
}
