package domain

import (
	"fmt"
	"time"
)

type Address struct {
	Id               int64     `json:"id"`
	Address1         string    `json:"address1"`
	Address2         string    `json:"address2"`
	AlternativePhone string    `json:"alternative_phone"`
	City             string    `json:"city"`
	Company          string    `json:"company"`
	CountryId        int64     `json:"country_id"`
	Firstname        string    `json:"firstname"`
	FullName         string    `json:"full_name"`
	Lastname         string    `json:"lastname"`
	Phone            string    `json:"phone"`
	StateId          int64     `json:"state_id"`
	Zipcode          string    `json:"zipcode"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`

	//Computed
	StateName string `json:"state_name"`
	StateText string `json:"state_text"` //TODO:Check implementation in interactor

	// Associations
	Country *Country `json:"country"`
	State   *State   `json:"state"`
}

func (this *Address) AfterFind() (err error) {
	this.FullName = fmt.Sprintf("%s %s", this.Firstname, this.Lastname)
	return
}

func (this Address) TableName() string {
	return "spree_addresses"
}
